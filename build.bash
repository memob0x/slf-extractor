#!/usr/bin/env bash
# thanks to https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures

package=github.com/memob0x/slf-exporter

package_split=(${package//\// })
package_name=${package_split[-1]}

# go tool dist list
platforms=("windows/amd64" "windows/386" "linux/386" "linux/amd64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$package_name'-'$GOOS'-'$GOARCH
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	

	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done