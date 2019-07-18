#!/bin/sh
# Generates a script from a given schema zip file, finds the names of the
# objects, then generates go files based on the provided
# generate_from_schema.py tool

# Find the schema document name by going here: https://www.dmtf.org/standards/redfish
# Inspect the url for the schema you want - for example, the 2019.1 update document is "DSP8010" on this page, the zip file is:
schemadoc=DSP8010_2019.1_0.zip
#
echo "Fetching schema document $schemadoc"
curl -G -L https://www.dmtf.org/sites/default/files/standards/documents/$schemadoc > wrk.zip
unzip -v wrk.zip | grep / | grep "json-schema" | cut -c 59- | cut -d "/" -f 3 |  grep . | grep -v "Collection" | grep -v "Schedule" | grep -v "redfish-schema" | grep -v ".v" | grep -v "odata" | grep -v "Protocol" | cut -d . -f 1 >f1.txt
rm wrk.zip
# Now we're ready to populate the script file
echo "mkdir gofiles" > gen_go_source.sh
lam -s "echo \"Getting object \\\"" f1.txt -s "\\\" and generating go source...\";python3 generate_from_schema.py " f1.txt -s " > gofiles/" f1.txt -s ".go " >> gen_go_source.sh
echo "echo \"Go source has been created in ./gofiles/\"" >> gen_go_source.sh
# Done. Finish cleaning up
rm f1.txt
#Execute the script
source ./gen_go_source.sh
