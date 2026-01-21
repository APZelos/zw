package shell

// FishIntegration returns the fish shell integration script.
func FishIntegration() string {
	return `# zw shell integration for fish
function zw
    if test (count $argv) -gt 0
        switch $argv[1]
            case init version --help -h
                command zw $argv
                return
        end
    end

    set -l result (command zw --query $argv)
    if test -n "$result"
        cd $result
    end
end
`
}
