#!/usr/bin/env bash
# thanks to https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures

function has_param() {
    local terms="$1"
    shift

    for term in $terms; do
        for arg; do
            if [[ $arg == "$term" ]]; then
                echo "yes"
            fi
        done
    done
}

package=github.com/memob0x/slf-exporter

package_split=(${package//\// })
package_name=${package_split[-1]}
 
# go tool dist list
platforms=("linux/amd64" "darwin/amd64" "windows/amd64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	
	GOOS=${platform_split[0]}
	
	GOARCH=${platform_split[1]}

	if [[ -n $(has_param "-h --with-gui" "$@") ]]; then
		fyne-cross $GOOS --pull -arch=$GOARCH -app-id slf-exporter
	fi

	output_name=$package_name'-cli-'$GOOS'-'$GOARCH

	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi

	env GOOS=$GOOS GOARCH=$GOARCH go build -tags=cli -o $output_name $package

	if [ $? -ne 0 ]; then
		echo 'An error has occurred! Aborting the script execution...'

		exit 1
	fi
done