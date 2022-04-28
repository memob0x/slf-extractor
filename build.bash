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

package=github.com/memob0x/slf-exporter

package_split=(${package//\// })
package_name=${package_split[-1]}
 
# go tool dist list
platforms=("linux/amd64" "darwin/amd64" "windows/amd64")

# cleanup
if compgen -G "fyne-cross" > /dev/null;
then
	echo "clean previous fyne build folder"

	rm -r fyne-cross
fi

if compgen -G "slf-exporter-*" > /dev/null;
then
	echo "clean previous build files"

	rm slf-exporter*
fi 

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	
	GOOS=${platform_split[0]}
	
	GOARCH=${platform_split[1]}

	# FIXME: darwin (osx) gui build needs special passages https://github.com/fyne-io/fyne-cross#build_darwin_image
	if [[ -n $(has_param "-h --with-gui" "$@") ]];
	then
 		../../../../bin/fyne-cross $GOOS --pull -arch=$GOARCH -app-id slf-exporter -tags=gui
	fi

	output_name=$package_name'-cli-'$GOOS'-'$GOARCH

	if [ $GOOS = "windows" ];
	then
		output_name+='.exe'
	fi

	env GOOS=$GOOS GOARCH=$GOARCH go build -tags=cli -o $output_name $package

	if [ $? -ne 0 ];
	then
		echo 'An error has occurred! Aborting the script execution...'

		exit 1
	fi
done