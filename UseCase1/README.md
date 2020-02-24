# Use Case 1 
The following use case is written to achieve the following;
- Stand alone Pod deployed in kubernetes.
- Pod will run a node/go application (container 1)
- The app will write to a file (file.txt) every 5 mins [datetimestamp] and output of GET "http://worldtimeapi.org/api/timezone/Europe/London.txt" parsing only the datetime field
- The output directory will be /tmp/data
- Pod will run a second container which will dump the contents of the file from container 1 into STDOUT
-  emptydir volume need to be created and mounted on the pod at /tmp/data and /tmp/output
- Use the final images as UBI image for this 
