#!/bin/sh
# Generates a subordinate script (gen_go_source.sh) from a given schema zip
# file, finds the names of the objects, then generates go files based on
# the provided generate_from_schema.py tool and accompanying source.tmpl file.
# Leaves the gen_go_source.sh in place for modification if needed and can be run
# independently of this script to regenerate files.

# Find the schema document name by going here:
#                                       https://www.dmtf.org/standards/redfish
# Inspect the url for the schema you want - for example, the 2019.1 update
# document is "DSP8010" on this page, the zip file is:
schemadoc=DSP8010_2019.1_0.zip
#
echo "Fetching schema document $schemadoc"
# We only use the zip file to get the object names. It's a little easier to
# manage.
curl -G -L https://www.dmtf.org/sites/default/files/standards/documents/$schemadoc > wrk.zip
# Generate the object list. No elegant, but it works well enough for now.
#
# SerialInterface and Switch still have some identifier issues to be worked out
# Collection files, Schedule, Protocol are "common" and included differently
# redfish-schema, odata and all the versions of each object aren't needed
#
# General process is get a list of the json-schema objects from the zip, drop
# things we don't need/want, and clip the column we want generating a file of
# object names we can use later.
unzip -v wrk.zip | grep / | grep "json-schema" | cut -c 59- | cut -d "/" -f 3 |  grep . | grep -v "Collection" | grep -v "Schedule" | grep -v "redfish-schema" | grep -v ".v" | grep -v "odata" | grep -v "Protocol" | cut -d . -f 1 >f1.txt
# f1.txt now has our list of objects
# don't need the document zip file any more...
rm wrk.zip
# Now we're ready to populate the script file that will actually generate the
# code files.
echo "mkdir gofiles" > gen_go_source.sh
lam -s "echo \"Getting object \\\"" f1.txt -s "\\\" and generating go source...\";python3 generate_from_schema.py " f1.txt -s " > gofiles/" f1.txt -s ".go " >> gen_go_source.sh
echo "echo \"Go source has been created in ./gofiles/\"" >> gen_go_source.sh
echo "echo \"Executing go fmt * ...\"" >> gen_go_source.sh
echo "go fmt ./gofiles/*.go" >> gen_go_source.sh
echo "echo \"Ready for manual cleanup.\"" >> gen_go_source.sh
# Done. Finish cleaning up
rm f1.txt
#Execute the script
sh ./gen_go_source.sh
