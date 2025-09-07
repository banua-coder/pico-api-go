#!/usr/bin/env ruby
# frozen_string_literal: true

require 'date'
require 'optparse'

##
# Automatic Changelog Generator (Ruby Version)
#
# This script automatically generates changelog entries based on git commits
# following the Keep a Changelog format and Semantic Versioning principles.
#
# This Ruby implementation offers several advantages over the bash version:
# - Better error handling and validation with structured exception handling
# - More robust parsing of conventional commit messages
# - Cleaner code organization with object-oriented design
# - More reliable text processing without shell escaping issues
# - Better support for complex commit message parsing
# - More maintainable and testable code structure
# - Cross-platform compatibility (Ruby vs. bash-specific features)
#
# Features:
# - Only runs from release branches (release/vx.x.x)
# - Categorizes commits by conventional commit types
# - Determines semantic version increment automatically
# - Updates CHANGELOG.md with proper formatting
# - Robust error handling and validation
# - Dry-run mode for previewing changes
# - Force mode for bypassing uncommitted changes check
#
# Usage:
#   ruby generate-changelog.rb [options]
#
# Author: Auto-generated for PICO API Go project
# License: Same as project license
class ChangelogGenerator
  # Conventional commit types and their changelog categories
  COMMIT_CATEGORIES = {
    'feat' => { category: 'Added', breaking: false },
    'fix' => { category: 'Fixed', breaking: false },
    'docs' => { category: 'Documentation', breaking: false },
    'style' => { category: 'Style', breaking: false },
    'refactor' => { category: 'Changed', breaking: false },
    'perf' => { category: 'Performance', breaking: false },
    'test' => { category: 'Tests', breaking: false },
    'chore' => { category: 'Maintenance', breaking: false },
    'ci' => { category: 'CI/CD', breaking: false },
    'build' => { category: 'Build', breaking: false },
    'revert' => { category: 'Reverted', breaking: false }
  }.freeze

  # Release branch pattern
  RELEASE_BRANCH_PATTERN = /^release\/v(\d+)\.(\d+)\.(\d+)$/

  # Changelog file path
  CHANGELOG_PATH = 'CHANGELOG.md'

  attr_reader :options, :current_branch, :version_info

  ##
  # Initialize the changelog generator
  #
  # @param options [Hash] Configuration options
  def initialize(options = {})
    @options = default_options.merge(options)
    @current_branch = get_current_branch
    @version_info = parse_version_from_branch
    validate_environment!
  end

  ##
  # Generate the changelog entry
  #
  # @return [Boolean] true if successful, false otherwise
  def generate!
    puts "🚀 Generating changelog for version #{version_string}..."
    
    commits = fetch_commits_since_last_release
    if commits.empty?
      puts "⚠️  No commits found since last release. Nothing to generate."
      return false
    end

    categorized_commits = categorize_commits(commits)
    version_bump = determine_version_bump(commits)
    
    if options[:dry_run]
      preview_changelog(categorized_commits, version_bump)
    else
      update_changelog(categorized_commits, version_bump)
      puts "✅ Changelog updated successfully!"
    end

    true
  rescue StandardError => e
    puts "❌ Error generating changelog: #{e.message}"
    puts e.backtrace if options[:debug]
    false
  end

  private

  ##
  # Default configuration options
  #
  # @return [Hash] Default options
  def default_options
    {
      dry_run: false,
      debug: false,
      force: false,
      output_format: :markdown
    }
  end

  ##
  # Get the current git branch name
  #
  # @return [String] Current branch name
  # @raise [RuntimeError] if not in a git repository
  def get_current_branch
    branch = `git branch --show-current`.strip
    raise "Not in a git repository" if branch.empty?
    branch
  end

  ##
  # Parse version information from the current branch name
  #
  # @return [Hash] Version components (major, minor, patch)
  # @raise [RuntimeError] if not on a valid release branch
  def parse_version_from_branch
    match = current_branch.match(RELEASE_BRANCH_PATTERN)
    raise "Not on a release branch. Expected format: release/vX.Y.Z" unless match

    {
      major: match[1].to_i,
      minor: match[2].to_i,
      patch: match[3].to_i
    }
  end

  ##
  # Get the current version string
  #
  # @return [String] Version string (e.g., "v1.2.3")
  def version_string
    "v#{version_info[:major]}.#{version_info[:minor]}.#{version_info[:patch]}"
  end

  ##
  # Validate the environment before proceeding
  #
  # @raise [RuntimeError] if environment is invalid
  def validate_environment!
    # Check if git is available
    system('git --version > /dev/null 2>&1') || raise("Git is not installed or not in PATH")

    # Check if we're in a git repository
    system('git rev-parse --git-dir > /dev/null 2>&1') || raise("Not in a git repository")

    # Check if CHANGELOG.md exists
    File.exist?(CHANGELOG_PATH) || raise("#{CHANGELOG_PATH} not found")

    # Warn if there are uncommitted changes
    if has_uncommitted_changes? && !options[:force]
      raise "There are uncommitted changes. Use --force to proceed anyway."
    end

    puts "✅ Environment validation passed"
  end

  ##
  # Check if there are uncommitted changes
  #
  # @return [Boolean] true if there are uncommitted changes
  def has_uncommitted_changes?
    !system('git diff --quiet && git diff --cached --quiet')
  end

  ##
  # Fetch commits since the last release tag
  #
  # @return [Array<Hash>] Array of commit information
  def fetch_commits_since_last_release
    last_tag = get_last_release_tag
    range = last_tag ? "#{last_tag}..HEAD" : "HEAD"
    
    puts "📋 Fetching commits since #{last_tag || 'beginning'}..."
    
    commit_format = '%H|%s|%b|%an|%ae|%ad'
    commits_output = `git log #{range} --pretty=format:"#{commit_format}" --date=iso`
    
    commits_output.split("\n").map do |line|
      parts = line.split('|', 6)
      next if parts.length < 6

      {
        hash: parts[0],
        subject: parts[1],
        body: parts[2],
        author_name: parts[3],
        author_email: parts[4],
        date: parts[5]
      }
    end.compact
  end

  ##
  # Get the last release tag
  #
  # @return [String, nil] Last release tag or nil if none exists
  def get_last_release_tag
    tags = `git tag -l --sort=-version:refname`.split("\n")
    tags.find { |tag| tag.match?(/^v\d+\.\d+\.\d+$/) }
  end

  ##
  # Categorize commits by their conventional commit type
  #
  # @param commits [Array<Hash>] Array of commit information
  # @return [Hash] Commits grouped by category
  def categorize_commits(commits)
    categories = Hash.new { |h, k| h[k] = [] }
    
    commits.each do |commit|
      type, scope, description, breaking = parse_conventional_commit(commit[:subject])
      
      # Determine category
      category_info = COMMIT_CATEGORIES[type] || { category: 'Other', breaking: false }
      category = breaking ? 'Breaking Changes' : category_info[:category]
      
      # Skip certain types if configured
      next if should_skip_commit?(type, commit)
      
      categories[category] << {
        type: type,
        scope: scope,
        description: description,
        breaking: breaking,
        commit: commit
      }
    end
    
    # Remove empty categories and sort
    categories.reject { |_, commits| commits.empty? }
             .sort_by { |category, _| category_priority(category) }
             .to_h
  end

  ##
  # Parse a conventional commit message
  #
  # @param subject [String] Commit subject line
  # @return [Array] [type, scope, description, breaking]
  def parse_conventional_commit(subject)
    # Match conventional commit format: type(scope): description
    match = subject.match(/^(\w+)(?:\(([^)]+)\))?(!)?: (.+)$/)
    
    if match
      type = match[1].downcase
      scope = match[2]
      breaking_marker = match[3] == '!'
      description = match[4]
      
      # Check for BREAKING CHANGE in description
      breaking = breaking_marker || description.include?('BREAKING CHANGE')
      
      [type, scope, description, breaking]
    else
      # Fallback for non-conventional commits
      ['other', nil, subject, false]
    end
  end

  ##
  # Determine if a commit should be skipped
  #
  # @param type [String] Commit type
  # @param commit [Hash] Commit information
  # @return [Boolean] true if commit should be skipped
  def should_skip_commit?(type, commit)
    # Skip merge commits
    return true if commit[:subject].start_with?('Merge ')
    
    # Skip certain types if configured
    skip_types = options[:skip_types] || []
    skip_types.include?(type)
  end

  ##
  # Get priority for category ordering
  #
  # @param category [String] Category name
  # @return [Integer] Priority (lower = higher priority)
  def category_priority(category)
    priorities = {
      'Breaking Changes' => 1,
      'Added' => 2,
      'Changed' => 3,
      'Fixed' => 4,
      'Deprecated' => 5,
      'Removed' => 6,
      'Security' => 7,
      'Performance' => 8,
      'Documentation' => 9,
      'Tests' => 10,
      'CI/CD' => 11,
      'Build' => 12,
      'Maintenance' => 13,
      'Other' => 14
    }
    priorities[category] || 99
  end

  ##
  # Determine the version bump type based on commits
  #
  # @param commits [Array<Hash>] Array of commits
  # @return [Symbol] :major, :minor, or :patch
  def determine_version_bump(commits)
    has_breaking = commits.any? do |commit|
      _, _, _, breaking = parse_conventional_commit(commit[:subject])
      breaking || commit[:body].include?('BREAKING CHANGE')
    end
    
    return :major if has_breaking
    
    has_features = commits.any? do |commit|
      type, _, _, _ = parse_conventional_commit(commit[:subject])
      type == 'feat'
    end
    
    has_features ? :minor : :patch
  end

  ##
  # Preview the changelog without writing to file
  #
  # @param categorized_commits [Hash] Commits grouped by category
  # @param version_bump [Symbol] Type of version bump
  def preview_changelog(categorized_commits, version_bump)
    puts "\n" + "="*50
    puts "CHANGELOG PREVIEW (#{version_bump.upcase} BUMP)"
    puts "="*50
    puts
    puts generate_changelog_content(categorized_commits)
  end

  ##
  # Update the CHANGELOG.md file
  #
  # @param categorized_commits [Hash] Commits grouped by category
  # @param version_bump [Symbol] Type of version bump
  def update_changelog(categorized_commits, version_bump)
    current_content = File.read(CHANGELOG_PATH)
    new_content = generate_changelog_content(categorized_commits)
    
    # Find the position to insert new content (after ## [Unreleased])
    unreleased_pattern = /^## \[Unreleased\]\s*\n/
    match = current_content.match(unreleased_pattern)
    
    unless match
      raise "Could not find [Unreleased] section in #{CHANGELOG_PATH}"
    end
    
    # Insert new release section after the unreleased section
    insertion_point = match.end(0)
    
    # Clear the unreleased section and add new release
    updated_content = current_content[0...insertion_point] +
                     "\n" +
                     new_content +
                     "\n" +
                     current_content[insertion_point..-1]
    
    # Update comparison links at the bottom
    updated_content = update_comparison_links(updated_content)
    
    # Write back to file
    File.write(CHANGELOG_PATH, updated_content)
  end

  ##
  # Generate changelog content for the new release
  #
  # @param categorized_commits [Hash] Commits grouped by category
  # @return [String] Formatted changelog content
  def generate_changelog_content(categorized_commits)
    content = []
    content << "## [#{version_string}] - #{Date.today.strftime('%Y-%m-%d')}"
    content << ""
    
    categorized_commits.each do |category, commits|
      content << "### #{category}"
      content << ""
      
      commits.each do |commit_info|
        line = format_changelog_line(commit_info)
        content << line if line
      end
      
      content << ""
    end
    
    content.join("\n")
  end

  ##
  # Format a single changelog line
  #
  # @param commit_info [Hash] Commit information
  # @return [String, nil] Formatted changelog line
  def format_changelog_line(commit_info)
    description = commit_info[:description]
    scope = commit_info[:scope]
    
    # Format: "- description (scope if present)"
    line = "- #{description.capitalize}"
    line += " (#{scope})" if scope && !scope.empty?
    
    # Add breaking change marker
    line = "- **BREAKING**: #{description}" if commit_info[:breaking]
    
    line
  end

  ##
  # Update comparison links at the bottom of the changelog
  #
  # @param content [String] Current changelog content
  # @return [String] Updated changelog content with new comparison links
  def update_comparison_links(content)
    # This would need to be customized based on your repository URL structure
    # For now, we'll leave the existing links unchanged
    content
  end
end

##
# Command line interface
class CLI
  def self.run(args = ARGV)
    options = {}
    
    parser = OptionParser.new do |opts|
      opts.banner = "Usage: #{$0} [options]"
      opts.separator ""
      opts.separator "Automatic Changelog Generator"
      opts.separator ""
      opts.separator "This script generates changelog entries from git commits"
      opts.separator "following conventional commit format and Keep a Changelog style."
      opts.separator ""
      opts.separator "Requirements:"
      opts.separator "- Must be run from a release branch (release/vX.Y.Z)"
      opts.separator "- Git repository with existing tags"
      opts.separator "- CHANGELOG.md file with [Unreleased] section"
      opts.separator ""
      
      opts.on("-d", "--dry-run", "Preview changes without modifying files") do
        options[:dry_run] = true
      end
      
      opts.on("-f", "--force", "Proceed even with uncommitted changes") do
        options[:force] = true
      end
      
      opts.on("--debug", "Enable debug output") do
        options[:debug] = true
      end
      
      opts.on("-h", "--help", "Show this help message") do
        puts opts
        exit 0
      end
      
      opts.separator ""
      opts.separator "Examples:"
      opts.separator "  #{$0}                  # Generate changelog"
      opts.separator "  #{$0} --dry-run        # Preview without changes"
      opts.separator "  #{$0} --force          # Ignore uncommitted changes"
    end
    
    begin
      parser.parse!(args)
      
      generator = ChangelogGenerator.new(options)
      success = generator.generate!
      
      exit(success ? 0 : 1)
      
    rescue OptionParser::InvalidOption => e
      puts "Error: #{e.message}"
      puts parser
      exit 1
    rescue StandardError => e
      puts "Error: #{e.message}"
      exit 1
    end
  end
end

# Run the CLI if this file is executed directly
if __FILE__ == $0
  CLI.run
end