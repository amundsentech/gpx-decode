# GPX Decode

Plenty of libraries encode gpx, but few decode gpx, likely due to it being xml which can use pre-made parsers.  The advantage of a special decoder is gpx files have known schemas, are the default handheld GPS format, and can be commonly used as inputs for other uses such as viewing in 3D (eg. MineAR).  As gpx xml is verbose and repetative, ***GPX Decode*** makes decoding/unmarshaling them pretty useful... and parses the Extensions (extended attributes) that most other gpx decoders do not.


## Usage

````GPXDecode(gpxbuf, &gpx)````

Reads a gpx file into a gpx struct.  Decodes one/more point, line or polygon geometries, with/out z elevation values, with/out extended attributes.

## Note!

Extended attributes are by default type STRING.  This parser does _not_ create new struct fields encoded by type.

THEREFORE after decoding, a user of this object must create field types according to the contained data.
