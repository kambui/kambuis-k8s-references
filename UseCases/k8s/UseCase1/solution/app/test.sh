# Test Script for  write-to-files 

################################################################################
# Help                                                                         #
################################################################################
Help()
{
   # Display Help
   echo "This script test the write-to-file.go program."
   echo "It creates the image from the Dockerfile, then runs the image"
   echo "as a local container (using podman). Then tries to retrieve"
   echo "the generate files and does a pattern match to check output."
   echo "On success returns PASS or FAILED"
   echo ""
   echo "Syntax: scriptTemplate [-g|h|t|v|V]"
   echo ""
   echo "options:"
   echo "h     Print this Help."
   echo
}

# Get the options
while getopts ":h" option; do
   case $option in
      h) # display Help
         Help
         exit;;
     \?) # incorrect option
         echo "Error: Invalid option"
         exit;;
   esac
done

#podman build -t write-to-file:test

#podman run --name testrun write-to-file:test

#podman exec testrun cat /tmp/data/output
