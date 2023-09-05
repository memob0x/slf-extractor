#!/usr/bin/env bash
# thanks to https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures

function has_param() {
    local terms="$1"
    shift

    for term in $terms;
	do
        for arg;
		do
            if [[ $arg == "$term" ]];
			then
                echo "yes"
            fi
        done
    done
}

package=github.com/memob0x/slf-extractor

package_split=(${package//\// })
package_name=${package_split[-1]}

with_gui=$(has_param "-h --with-gui" "$@")
 
# go tool dist list
platforms=("linux/amd64" "darwin/amd64" "windows/amd64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	
	GOOS=${platform_split[0]}
	
	GOARCH=${platform_split[1]}

	output_name=$package_name

	if [ $GOOS = "windows" ];
	then
		output_name+='.exe'
	fi

	# NOTE: darwin gui build is not supported atm since xcode requirements are too strict
	# see https://github.com/fyne-io/fyne-cross/blob/3b7d42345c647393361e9eea2335921b33eaabcf/README.md#build-the-docker-image-for-osxdarwinapple-cross-compiling
	if [[ "$with_gui" == "yes" ]] && [[ $GOOS != "darwin" ]];
	then
 		fyne-cross $GOOS --pull -arch $GOARCH -tags gui -output $output_name

		mkdir -p dist/gui/$GOOS-$GOARCH

		cp -r fyne-cross/bin/$GOOS-$GOARCH dist/gui/$GOOS-$GOARCH

		rm -rf fyne-cross
	fi

	env GOOS=$GOOS GOARCH=$GOARCH go build -tags cli -o dist/cli/$GOOS-$GOARCH/$output_name_cli $package

	if [ $? -ne 0 ];
	then
		echo 'An error has occurred! Aborting the script execution...'

		exit 1
	fi
done
