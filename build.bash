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
 
# go tool dist list
platforms=("linux/amd64" "darwin/amd64" "windows/amd64")

# cleanup
if compgen -G "fyne-cross" > /dev/null;
then
	echo "clean previous gui build folder"

	rm -r fyne-cross
fi

if compgen -G "fyne-build-cli" > /dev/null;
then
	echo "clean previous cli build files"

	rm -r fyne-build-cli
fi

mkdir fyne-build-cli

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	
	GOOS=${platform_split[0]}
	
	GOARCH=${platform_split[1]}

	output_name_gui=$package_name'-gui-'$GOOS'-'$GOARCH

	# FIXME: darwin (osx) gui build needs special passages https://github.com/fyne-io/fyne-cross#build_darwin_image
	if [[ -n $(has_param "-h --with-gui" "$@") ]];
	then
 		fyne-cross $GOOS --pull -arch=$GOARCH -app-id slf-extractor -tags=gui -icon assets/icon.png -output $output_name_gui
	fi

	output_name_cli=$package_name'-cli-'$GOOS'-'$GOARCH

	if [ $GOOS = "windows" ];
	then
		output_name_cli+='.exe'
	fi

	env GOOS=$GOOS GOARCH=$GOARCH go build -tags=cli -o fyne-build-cli/$output_name_cli $package

	if [ $? -ne 0 ];
	then
		echo 'An error has occurred! Aborting the script execution...'

		exit 1
	fi
done