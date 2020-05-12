# Exiftool

## Getting started
Find the different from DateTimeOriginal and GPSDateTime
```bash
$ ./main -i ~/Pictures/GoPro/wrongdate/2016-01-01 diff
```


If not GPSDateTime available but Date is know
```bash
$ ./main -i ~/Pictures/GoPro/wrongdate/2016-01-06 -offset -6 check 20200506T2025
> diff minutes:  2278942
```

## Steps
1. $ ./main -i ~/Pictures/GoPro/wrongdate/2016-01-06 -offset -6 check 20200506T2025
    > diff minutes:  2278942

2. $ ./main -i ~/Pictures/GoPro/wrongdate/2016-01-06 -offset 0 date 2278942
3. $ exiftool '-Directory<CreateDate' -d %Y-%m-%d -r .

## Importand
Always double check with
```bash
exiftool -alldates -gpsdatetime .
```

If something goes wrong, revert it
```bash
exiftool -restore_original .
```
