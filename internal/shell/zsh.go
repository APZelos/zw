package shell

// ZshIntegration returns the zsh shell integration script.
func ZshIntegration() string {
	return `# zw shell integration for zsh
zw() {
    if [[ "$1" == "init" || "$1" == "version" || "$1" == "--help" || "$1" == "-h" ]]; then
        command zw "$@"
        return
    fi

    local result
    result=$(command zw --query "$@")
    if [[ -n "$result" ]]; then
        cd "$result" || return 1
    fi
}
`
}
