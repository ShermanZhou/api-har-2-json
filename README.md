# api-har-2-json

### Description

extract api's json data to individual json files.
(1) sequenced from first api; this reserved original sequency, so if same api are extracted, you can choose which to keep.
(2) GET, POST, PUT calls are extracted, DELETE calls are not extracted.

### Usage:
(1) create a folder foo
(2) copy the har file in the foo/
(3) run the executable to 
 $ api-har-2-json  -har  bar-har-file-name.har
 
 You will have extracted json:  
 
 Naming convention:
 - api's path 's 'slash' is translated to dash '-' into json file's n
