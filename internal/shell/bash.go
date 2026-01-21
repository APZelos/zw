package shell

// BashIntegration returns the bash shell integration script.
func BashIntegration() string {
	return `# zw shell integration for bash
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
