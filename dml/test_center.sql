INSERT INTO testing_centers (center_name, address, days_open, time_open, time_closed, website, need_appointment)
VALUES
(
'Hey Denver',
'1720 Pearl Street',
'Monday, Tuesday, Wendesday, Thursday, Friday, Saturday, Sunday', 0230,
0730,
'http://heydenver.org/',
false
),

(
'Denver Health',
'1720 hospital street',
'Monday, Tuesday, Wendesday, Thursday, Friday, Saturday, Sunday', 0230,
0730,
'http://denverhealth.com/',
false
)


http POST http://localhost:8000/api/v1/centers \
    Center_name=A place \
    Address="1111 place st" married:=false hobbies:='["http", "pies"]' \  # Raw JSON
    description=@about-john.txt \   # Embed text file
    bookmarks:=@bookmarks.json


curl -i -X POST -H "Content-Type: application/json" -d "{\"Center_name\": \"A place\", \"Address\":\"somewhere\", \"Days_open\": \"Monday, Tuesday\", \"Time_open\":\6\, \"Time_closed\":\10\, \"Website\": \"www.aplace.com\", \"Need_appointment\": \false\ }" http://localhost:8000/api/v1/centers
