window.BENCHMARK_DATA = {
  "lastUpdate": 1731928167202,
  "repoUrl": "https://github.com/nobl9/govy",
  "entries": {
    "Govy Go Benchmark": [
      {
        "commit": {
          "author": {
            "email": "48822818+nieomylnieja@users.noreply.github.com",
            "name": "Mateusz Hawrus",
            "username": "nieomylnieja"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "ddb940d6e370cacf2017a35db7057829a83fa4a5",
          "message": "chore: Split test workflows for pushes to main (#45)\n\n## Summary\r\n\r\nCurrently we're adding `if` statements to assert a GitHub event is equal\r\nto `push` to run some of the push-only actions.\r\nThis is not ideal as we're introducing misconfiguration potential and\r\nelevate the job's permissions.\r\nFurthermore, it seems that it might be easier to configure [benchmark\r\naction](https://github.com/benchmark-action/github-action-benchmark)\r\nseparately for publishing GitHub Pages and for checking the PR.",
          "timestamp": "2024-11-05T13:21:02+01:00",
          "tree_id": "5fc56448451f18b167fdde40b6c5821f2d58cb9e",
          "url": "https://github.com/nobl9/govy/commit/ddb940d6e370cacf2017a35db7057829a83fa4a5"
        },
        "date": 1730809399508,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 672.1,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1791199 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 672.1,
            "unit": "ns/op",
            "extra": "1791199 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1791199 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1791199 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 778,
            "unit": "ns/op\t     240 B/op\t       6 allocs/op",
            "extra": "1541877 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 778,
            "unit": "ns/op",
            "extra": "1541877 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "1541877 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1541877 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 814.1,
            "unit": "ns/op\t     400 B/op\t      10 allocs/op",
            "extra": "1472438 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 814.1,
            "unit": "ns/op",
            "extra": "1472438 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 400,
            "unit": "B/op",
            "extra": "1472438 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1472438 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 743,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1610506 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 743,
            "unit": "ns/op",
            "extra": "1610506 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1610506 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1610506 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 802.6,
            "unit": "ns/op\t     376 B/op\t      10 allocs/op",
            "extra": "1407355 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 802.6,
            "unit": "ns/op",
            "extra": "1407355 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 376,
            "unit": "B/op",
            "extra": "1407355 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1407355 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 767.7,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1563001 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 767.7,
            "unit": "ns/op",
            "extra": "1563001 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1563001 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1563001 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1197,
            "unit": "ns/op\t     520 B/op\t      18 allocs/op",
            "extra": "900819 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1197,
            "unit": "ns/op",
            "extra": "900819 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "900819 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "900819 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 172.2,
            "unit": "ns/op\t     144 B/op\t       4 allocs/op",
            "extra": "7063042 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 172.2,
            "unit": "ns/op",
            "extra": "7063042 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 144,
            "unit": "B/op",
            "extra": "7063042 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "7063042 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1317,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "848688 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1317,
            "unit": "ns/op",
            "extra": "848688 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "848688 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "848688 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength",
            "value": 1016,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - ns/op",
            "value": 1016,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength",
            "value": 1018,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "991920 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1018,
            "unit": "ns/op",
            "extra": "991920 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "991920 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "991920 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength",
            "value": 1234,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "888162 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1234,
            "unit": "ns/op",
            "extra": "888162 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "888162 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "888162 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 1025,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 1025,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength",
            "value": 1051,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - ns/op",
            "value": 1051,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength",
            "value": 1054,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1054,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - B/op",
            "value": 560,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength",
            "value": 1025,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - ns/op",
            "value": 1025,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength",
            "value": 1079,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "983133 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - ns/op",
            "value": 1079,
            "unit": "ns/op",
            "extra": "983133 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "983133 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "983133 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf",
            "value": 1105,
            "unit": "ns/op\t     536 B/op\t      22 allocs/op",
            "extra": "958146 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1105,
            "unit": "ns/op",
            "extra": "958146 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 536,
            "unit": "B/op",
            "extra": "958146 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "958146 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 7148,
            "unit": "ns/op\t    3168 B/op\t      98 allocs/op",
            "extra": "165039 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 7148,
            "unit": "ns/op",
            "extra": "165039 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3168,
            "unit": "B/op",
            "extra": "165039 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "165039 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2614,
            "unit": "ns/op\t    1064 B/op\t      32 allocs/op",
            "extra": "437707 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2614,
            "unit": "ns/op",
            "extra": "437707 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1064,
            "unit": "B/op",
            "extra": "437707 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "437707 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1052,
            "unit": "ns/op\t     704 B/op\t      23 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1052,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - B/op",
            "value": 704,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty",
            "value": 180.2,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "6321806 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 180.2,
            "unit": "ns/op",
            "extra": "6321806 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "6321806 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6321806 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 1492,
            "unit": "ns/op\t     644 B/op\t      20 allocs/op",
            "extra": "737235 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 1492,
            "unit": "ns/op",
            "extra": "737235 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 644,
            "unit": "B/op",
            "extra": "737235 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "737235 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 1496,
            "unit": "ns/op\t     693 B/op\t      20 allocs/op",
            "extra": "742516 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 1496,
            "unit": "ns/op",
            "extra": "742516 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 693,
            "unit": "B/op",
            "extra": "742516 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "742516 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 15135,
            "unit": "ns/op\t    5640 B/op\t     154 allocs/op",
            "extra": "77636 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 15135,
            "unit": "ns/op",
            "extra": "77636 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 5640,
            "unit": "B/op",
            "extra": "77636 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 154,
            "unit": "allocs/op",
            "extra": "77636 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4363,
            "unit": "ns/op\t    1552 B/op\t      41 allocs/op",
            "extra": "267367 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4363,
            "unit": "ns/op",
            "extra": "267367 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1552,
            "unit": "B/op",
            "extra": "267367 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "267367 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 16127,
            "unit": "ns/op\t   14085 B/op\t     217 allocs/op",
            "extra": "74044 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 16127,
            "unit": "ns/op",
            "extra": "74044 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 14085,
            "unit": "B/op",
            "extra": "74044 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "74044 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 8593,
            "unit": "ns/op\t    3408 B/op\t     138 allocs/op",
            "extra": "137911 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 8593,
            "unit": "ns/op",
            "extra": "137911 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3408,
            "unit": "B/op",
            "extra": "137911 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "137911 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 8849,
            "unit": "ns/op\t    5904 B/op\t      54 allocs/op",
            "extra": "134198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 8849,
            "unit": "ns/op",
            "extra": "134198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5904,
            "unit": "B/op",
            "extra": "134198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "134198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1057,
            "unit": "ns/op\t     752 B/op\t      25 allocs/op",
            "extra": "987164 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1057,
            "unit": "ns/op",
            "extra": "987164 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "987164 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "987164 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP",
            "value": 1599,
            "unit": "ns/op\t     824 B/op\t      31 allocs/op",
            "extra": "686692 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1599,
            "unit": "ns/op",
            "extra": "686692 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 824,
            "unit": "B/op",
            "extra": "686692 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "686692 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1700,
            "unit": "ns/op\t     896 B/op\t      32 allocs/op",
            "extra": "662527 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1700,
            "unit": "ns/op",
            "extra": "662527 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "662527 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "662527 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1826,
            "unit": "ns/op\t    1056 B/op\t      36 allocs/op",
            "extra": "554678 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1826,
            "unit": "ns/op",
            "extra": "554678 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "554678 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "554678 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 3003,
            "unit": "ns/op\t    1576 B/op\t      63 allocs/op",
            "extra": "380036 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 3003,
            "unit": "ns/op",
            "extra": "380036 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "380036 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "380036 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5599,
            "unit": "ns/op\t    3048 B/op\t     118 allocs/op",
            "extra": "207870 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5599,
            "unit": "ns/op",
            "extra": "207870 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 3048,
            "unit": "B/op",
            "extra": "207870 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "207870 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3544,
            "unit": "ns/op\t    2056 B/op\t      75 allocs/op",
            "extra": "328495 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3544,
            "unit": "ns/op",
            "extra": "328495 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 2056,
            "unit": "B/op",
            "extra": "328495 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "328495 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1173,
            "unit": "ns/op\t     616 B/op\t      23 allocs/op",
            "extra": "927267 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1173,
            "unit": "ns/op",
            "extra": "927267 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 616,
            "unit": "B/op",
            "extra": "927267 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "927267 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2338,
            "unit": "ns/op\t    1448 B/op\t      44 allocs/op",
            "extra": "487446 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2338,
            "unit": "ns/op",
            "extra": "487446 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "487446 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "487446 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2383,
            "unit": "ns/op\t    1576 B/op\t      46 allocs/op",
            "extra": "475176 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2383,
            "unit": "ns/op",
            "extra": "475176 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "475176 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "475176 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1302,
            "unit": "ns/op\t     784 B/op\t      24 allocs/op",
            "extra": "828422 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1302,
            "unit": "ns/op",
            "extra": "828422 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 784,
            "unit": "B/op",
            "extra": "828422 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "828422 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1298,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "829464 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1298,
            "unit": "ns/op",
            "extra": "829464 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "829464 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "829464 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1598,
            "unit": "ns/op\t     976 B/op\t      30 allocs/op",
            "extra": "693373 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1598,
            "unit": "ns/op",
            "extra": "693373 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "693373 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "693373 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 11688,
            "unit": "ns/op\t    3984 B/op\t      72 allocs/op",
            "extra": "101740 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 11688,
            "unit": "ns/op",
            "extra": "101740 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 3984,
            "unit": "B/op",
            "extra": "101740 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 72,
            "unit": "allocs/op",
            "extra": "101740 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 37298,
            "unit": "ns/op\t    7577 B/op\t      99 allocs/op",
            "extra": "32180 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 37298,
            "unit": "ns/op",
            "extra": "32180 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7577,
            "unit": "B/op",
            "extra": "32180 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "32180 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 37738,
            "unit": "ns/op\t    7834 B/op\t     108 allocs/op",
            "extra": "31725 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 37738,
            "unit": "ns/op",
            "extra": "31725 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7834,
            "unit": "B/op",
            "extra": "31725 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "31725 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 37333,
            "unit": "ns/op\t    7449 B/op\t     103 allocs/op",
            "extra": "32030 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 37333,
            "unit": "ns/op",
            "extra": "32030 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7449,
            "unit": "B/op",
            "extra": "32030 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "32030 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 16020,
            "unit": "ns/op\t    8721 B/op\t     217 allocs/op",
            "extra": "74487 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 16020,
            "unit": "ns/op",
            "extra": "74487 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8721,
            "unit": "B/op",
            "extra": "74487 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "74487 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 156399419,
            "unit": "ns/op\t334255565 B/op\t  281343 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 156399419,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334255565,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281343,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 46804,
            "unit": "ns/op\t   29417 B/op\t     614 allocs/op",
            "extra": "25482 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 46804,
            "unit": "ns/op",
            "extra": "25482 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 29417,
            "unit": "B/op",
            "extra": "25482 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "25482 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4194,
            "unit": "ns/op\t    2054 B/op\t      58 allocs/op",
            "extra": "277071 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4194,
            "unit": "ns/op",
            "extra": "277071 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 2054,
            "unit": "B/op",
            "extra": "277071 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "277071 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1307,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "854250 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1307,
            "unit": "ns/op",
            "extra": "854250 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "854250 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "854250 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "48822818+nieomylnieja@users.noreply.github.com",
            "name": "Mateusz Hawrus",
            "username": "nieomylnieja"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "c2b1093171805dca3b72e3694998eff28304af69",
          "message": "chore: Update README.md (#46)\n\n## Summary\r\n\r\nUpdate both README.md and DEVELOPMENT.md with the coverage and\r\nbenchmarks sections.",
          "timestamp": "2024-11-05T14:23:34+01:00",
          "tree_id": "8cdce86abf78099b45f8687c0c77703a5b6540cf",
          "url": "https://github.com/nobl9/govy/commit/c2b1093171805dca3b72e3694998eff28304af69"
        },
        "date": 1730813149138,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 696.8,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1667223 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 696.8,
            "unit": "ns/op",
            "extra": "1667223 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1667223 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1667223 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 826.9,
            "unit": "ns/op\t     240 B/op\t       6 allocs/op",
            "extra": "1409127 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 826.9,
            "unit": "ns/op",
            "extra": "1409127 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "1409127 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1409127 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 853.6,
            "unit": "ns/op\t     400 B/op\t      10 allocs/op",
            "extra": "1426236 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 853.6,
            "unit": "ns/op",
            "extra": "1426236 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 400,
            "unit": "B/op",
            "extra": "1426236 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1426236 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 767.7,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1560285 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 767.7,
            "unit": "ns/op",
            "extra": "1560285 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1560285 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1560285 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 824.2,
            "unit": "ns/op\t     376 B/op\t      10 allocs/op",
            "extra": "1457751 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 824.2,
            "unit": "ns/op",
            "extra": "1457751 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 376,
            "unit": "B/op",
            "extra": "1457751 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1457751 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 783.7,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1500096 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 783.7,
            "unit": "ns/op",
            "extra": "1500096 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1500096 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1500096 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1233,
            "unit": "ns/op\t     520 B/op\t      18 allocs/op",
            "extra": "917168 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1233,
            "unit": "ns/op",
            "extra": "917168 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "917168 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "917168 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 176.2,
            "unit": "ns/op\t     144 B/op\t       4 allocs/op",
            "extra": "6786663 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 176.2,
            "unit": "ns/op",
            "extra": "6786663 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 144,
            "unit": "B/op",
            "extra": "6786663 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6786663 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1370,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "843876 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1370,
            "unit": "ns/op",
            "extra": "843876 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "843876 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "843876 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength",
            "value": 1052,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - ns/op",
            "value": 1052,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength",
            "value": 1057,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1057,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength",
            "value": 1279,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "858297 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1279,
            "unit": "ns/op",
            "extra": "858297 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "858297 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "858297 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 1055,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 1055,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength",
            "value": 1101,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "984256 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - ns/op",
            "value": 1101,
            "unit": "ns/op",
            "extra": "984256 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "984256 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "984256 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength",
            "value": 1097,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "961743 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1097,
            "unit": "ns/op",
            "extra": "961743 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - B/op",
            "value": 560,
            "unit": "B/op",
            "extra": "961743 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "961743 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength",
            "value": 1057,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - ns/op",
            "value": 1057,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength",
            "value": 1116,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "994158 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - ns/op",
            "value": 1116,
            "unit": "ns/op",
            "extra": "994158 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "994158 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "994158 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf",
            "value": 1138,
            "unit": "ns/op\t     536 B/op\t      22 allocs/op",
            "extra": "970252 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1138,
            "unit": "ns/op",
            "extra": "970252 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 536,
            "unit": "B/op",
            "extra": "970252 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "970252 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 7373,
            "unit": "ns/op\t    3168 B/op\t      98 allocs/op",
            "extra": "159813 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 7373,
            "unit": "ns/op",
            "extra": "159813 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3168,
            "unit": "B/op",
            "extra": "159813 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "159813 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2738,
            "unit": "ns/op\t    1064 B/op\t      32 allocs/op",
            "extra": "420727 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2738,
            "unit": "ns/op",
            "extra": "420727 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1064,
            "unit": "B/op",
            "extra": "420727 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "420727 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1085,
            "unit": "ns/op\t     704 B/op\t      23 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1085,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - B/op",
            "value": 704,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty",
            "value": 188.4,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "6397760 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 188.4,
            "unit": "ns/op",
            "extra": "6397760 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "6397760 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6397760 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 1535,
            "unit": "ns/op\t     644 B/op\t      20 allocs/op",
            "extra": "750741 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 1535,
            "unit": "ns/op",
            "extra": "750741 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 644,
            "unit": "B/op",
            "extra": "750741 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "750741 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 1535,
            "unit": "ns/op\t     693 B/op\t      20 allocs/op",
            "extra": "741579 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 1535,
            "unit": "ns/op",
            "extra": "741579 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 693,
            "unit": "B/op",
            "extra": "741579 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "741579 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 15482,
            "unit": "ns/op\t    5634 B/op\t     154 allocs/op",
            "extra": "77365 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 15482,
            "unit": "ns/op",
            "extra": "77365 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 5634,
            "unit": "B/op",
            "extra": "77365 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 154,
            "unit": "allocs/op",
            "extra": "77365 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4562,
            "unit": "ns/op\t    1552 B/op\t      41 allocs/op",
            "extra": "260070 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4562,
            "unit": "ns/op",
            "extra": "260070 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1552,
            "unit": "B/op",
            "extra": "260070 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "260070 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 16711,
            "unit": "ns/op\t   14085 B/op\t     217 allocs/op",
            "extra": "71352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 16711,
            "unit": "ns/op",
            "extra": "71352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 14085,
            "unit": "B/op",
            "extra": "71352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "71352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 8867,
            "unit": "ns/op\t    3408 B/op\t     138 allocs/op",
            "extra": "134462 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 8867,
            "unit": "ns/op",
            "extra": "134462 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3408,
            "unit": "B/op",
            "extra": "134462 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "134462 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 9157,
            "unit": "ns/op\t    5904 B/op\t      54 allocs/op",
            "extra": "130098 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 9157,
            "unit": "ns/op",
            "extra": "130098 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5904,
            "unit": "B/op",
            "extra": "130098 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "130098 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1085,
            "unit": "ns/op\t     752 B/op\t      25 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1085,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP",
            "value": 1643,
            "unit": "ns/op\t     824 B/op\t      31 allocs/op",
            "extra": "701288 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1643,
            "unit": "ns/op",
            "extra": "701288 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 824,
            "unit": "B/op",
            "extra": "701288 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "701288 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1760,
            "unit": "ns/op\t     896 B/op\t      32 allocs/op",
            "extra": "647350 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1760,
            "unit": "ns/op",
            "extra": "647350 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "647350 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "647350 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1891,
            "unit": "ns/op\t    1056 B/op\t      36 allocs/op",
            "extra": "610380 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1891,
            "unit": "ns/op",
            "extra": "610380 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "610380 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "610380 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 3107,
            "unit": "ns/op\t    1576 B/op\t      63 allocs/op",
            "extra": "372478 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 3107,
            "unit": "ns/op",
            "extra": "372478 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "372478 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "372478 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5786,
            "unit": "ns/op\t    3048 B/op\t     118 allocs/op",
            "extra": "201842 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5786,
            "unit": "ns/op",
            "extra": "201842 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 3048,
            "unit": "B/op",
            "extra": "201842 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "201842 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3665,
            "unit": "ns/op\t    2056 B/op\t      75 allocs/op",
            "extra": "319862 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3665,
            "unit": "ns/op",
            "extra": "319862 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 2056,
            "unit": "B/op",
            "extra": "319862 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "319862 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1215,
            "unit": "ns/op\t     616 B/op\t      23 allocs/op",
            "extra": "925156 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1215,
            "unit": "ns/op",
            "extra": "925156 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 616,
            "unit": "B/op",
            "extra": "925156 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "925156 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2412,
            "unit": "ns/op\t    1448 B/op\t      44 allocs/op",
            "extra": "486171 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2412,
            "unit": "ns/op",
            "extra": "486171 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "486171 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "486171 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2457,
            "unit": "ns/op\t    1576 B/op\t      46 allocs/op",
            "extra": "469231 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2457,
            "unit": "ns/op",
            "extra": "469231 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "469231 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "469231 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1347,
            "unit": "ns/op\t     784 B/op\t      24 allocs/op",
            "extra": "835191 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1347,
            "unit": "ns/op",
            "extra": "835191 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 784,
            "unit": "B/op",
            "extra": "835191 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "835191 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1336,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "825068 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1336,
            "unit": "ns/op",
            "extra": "825068 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "825068 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "825068 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1666,
            "unit": "ns/op\t     976 B/op\t      30 allocs/op",
            "extra": "696332 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1666,
            "unit": "ns/op",
            "extra": "696332 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "696332 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "696332 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 12052,
            "unit": "ns/op\t    3984 B/op\t      72 allocs/op",
            "extra": "99600 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 12052,
            "unit": "ns/op",
            "extra": "99600 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 3984,
            "unit": "B/op",
            "extra": "99600 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 72,
            "unit": "allocs/op",
            "extra": "99600 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 37798,
            "unit": "ns/op\t    7529 B/op\t      99 allocs/op",
            "extra": "31814 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 37798,
            "unit": "ns/op",
            "extra": "31814 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7529,
            "unit": "B/op",
            "extra": "31814 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "31814 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 38206,
            "unit": "ns/op\t    7833 B/op\t     108 allocs/op",
            "extra": "31240 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 38206,
            "unit": "ns/op",
            "extra": "31240 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7833,
            "unit": "B/op",
            "extra": "31240 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "31240 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 37816,
            "unit": "ns/op\t    7417 B/op\t     103 allocs/op",
            "extra": "31730 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 37816,
            "unit": "ns/op",
            "extra": "31730 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7417,
            "unit": "B/op",
            "extra": "31730 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "31730 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 16423,
            "unit": "ns/op\t    8721 B/op\t     217 allocs/op",
            "extra": "71264 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 16423,
            "unit": "ns/op",
            "extra": "71264 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8721,
            "unit": "B/op",
            "extra": "71264 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "71264 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 165194155,
            "unit": "ns/op\t334257220 B/op\t  281355 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 165194155,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334257220,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281355,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 47135,
            "unit": "ns/op\t   29417 B/op\t     614 allocs/op",
            "extra": "25195 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 47135,
            "unit": "ns/op",
            "extra": "25195 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 29417,
            "unit": "B/op",
            "extra": "25195 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "25195 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4210,
            "unit": "ns/op\t    2054 B/op\t      58 allocs/op",
            "extra": "277362 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4210,
            "unit": "ns/op",
            "extra": "277362 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 2054,
            "unit": "B/op",
            "extra": "277362 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "277362 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1312,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "831058 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1312,
            "unit": "ns/op",
            "extra": "831058 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "831058 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "831058 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "48822818+nieomylnieja@users.noreply.github.com",
            "name": "Mateusz Hawrus",
            "username": "nieomylnieja"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "78d44fd9640bce9bdf8b140c64af75262bb7921d",
          "message": "feat: Add AssertErrorContains (#47)\n\n## Motivation\r\n\r\nIn some cases we don't care about other errors and just want to check\r\nthat specific one is produced.\r\nTo achieve that we should add a helper which would work on a single\r\nerror, similar to how `assert.ErrorContains` works.\r\n\r\n## Release Notes\r\n\r\nAdded `govytest.AssertErrorContains` function which helps test govy\r\nrules by checking if a produced `govy.ValidatorError` contains specified\r\nerror.",
          "timestamp": "2024-11-06T17:56:47+01:00",
          "tree_id": "5cbc240b7b03d724cb52dea16c9395a864190d26",
          "url": "https://github.com/nobl9/govy/commit/78d44fd9640bce9bdf8b140c64af75262bb7921d"
        },
        "date": 1730912342843,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 679.4,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1771263 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 679.4,
            "unit": "ns/op",
            "extra": "1771263 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1771263 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1771263 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 782.1,
            "unit": "ns/op\t     240 B/op\t       6 allocs/op",
            "extra": "1536219 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 782.1,
            "unit": "ns/op",
            "extra": "1536219 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "1536219 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1536219 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 851.8,
            "unit": "ns/op\t     400 B/op\t      10 allocs/op",
            "extra": "1383440 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 851.8,
            "unit": "ns/op",
            "extra": "1383440 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 400,
            "unit": "B/op",
            "extra": "1383440 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1383440 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 780.5,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1539712 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 780.5,
            "unit": "ns/op",
            "extra": "1539712 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1539712 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1539712 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 839.9,
            "unit": "ns/op\t     376 B/op\t      10 allocs/op",
            "extra": "1505467 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 839.9,
            "unit": "ns/op",
            "extra": "1505467 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 376,
            "unit": "B/op",
            "extra": "1505467 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1505467 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 765,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1567134 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 765,
            "unit": "ns/op",
            "extra": "1567134 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1567134 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1567134 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1212,
            "unit": "ns/op\t     520 B/op\t      18 allocs/op",
            "extra": "903607 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1212,
            "unit": "ns/op",
            "extra": "903607 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "903607 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "903607 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 171.7,
            "unit": "ns/op\t     144 B/op\t       4 allocs/op",
            "extra": "7007996 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 171.7,
            "unit": "ns/op",
            "extra": "7007996 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 144,
            "unit": "B/op",
            "extra": "7007996 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "7007996 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1329,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "832572 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1329,
            "unit": "ns/op",
            "extra": "832572 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "832572 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "832572 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength",
            "value": 1030,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - ns/op",
            "value": 1030,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength",
            "value": 1051,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1051,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength",
            "value": 1249,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "845948 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1249,
            "unit": "ns/op",
            "extra": "845948 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "845948 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "845948 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 1048,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 1048,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength",
            "value": 1056,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - ns/op",
            "value": 1056,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength",
            "value": 1070,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1070,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - B/op",
            "value": 560,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength",
            "value": 1032,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - ns/op",
            "value": 1032,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength",
            "value": 1096,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "970747 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - ns/op",
            "value": 1096,
            "unit": "ns/op",
            "extra": "970747 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "970747 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "970747 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf",
            "value": 1124,
            "unit": "ns/op\t     536 B/op\t      22 allocs/op",
            "extra": "952652 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1124,
            "unit": "ns/op",
            "extra": "952652 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 536,
            "unit": "B/op",
            "extra": "952652 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "952652 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 7208,
            "unit": "ns/op\t    3168 B/op\t      98 allocs/op",
            "extra": "163333 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 7208,
            "unit": "ns/op",
            "extra": "163333 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3168,
            "unit": "B/op",
            "extra": "163333 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "163333 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2614,
            "unit": "ns/op\t    1064 B/op\t      32 allocs/op",
            "extra": "437641 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2614,
            "unit": "ns/op",
            "extra": "437641 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1064,
            "unit": "B/op",
            "extra": "437641 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "437641 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1054,
            "unit": "ns/op\t     704 B/op\t      23 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1054,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - B/op",
            "value": 704,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty",
            "value": 181.9,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "6601078 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 181.9,
            "unit": "ns/op",
            "extra": "6601078 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "6601078 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6601078 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 1501,
            "unit": "ns/op\t     644 B/op\t      20 allocs/op",
            "extra": "752256 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 1501,
            "unit": "ns/op",
            "extra": "752256 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 644,
            "unit": "B/op",
            "extra": "752256 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "752256 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 1500,
            "unit": "ns/op\t     692 B/op\t      20 allocs/op",
            "extra": "721512 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 1500,
            "unit": "ns/op",
            "extra": "721512 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 692,
            "unit": "B/op",
            "extra": "721512 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "721512 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 15200,
            "unit": "ns/op\t    5635 B/op\t     154 allocs/op",
            "extra": "77395 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 15200,
            "unit": "ns/op",
            "extra": "77395 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 5635,
            "unit": "B/op",
            "extra": "77395 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 154,
            "unit": "allocs/op",
            "extra": "77395 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4399,
            "unit": "ns/op\t    1552 B/op\t      41 allocs/op",
            "extra": "262909 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4399,
            "unit": "ns/op",
            "extra": "262909 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1552,
            "unit": "B/op",
            "extra": "262909 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "262909 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 16206,
            "unit": "ns/op\t   14085 B/op\t     217 allocs/op",
            "extra": "73677 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 16206,
            "unit": "ns/op",
            "extra": "73677 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 14085,
            "unit": "B/op",
            "extra": "73677 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "73677 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 8772,
            "unit": "ns/op\t    3408 B/op\t     138 allocs/op",
            "extra": "134048 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 8772,
            "unit": "ns/op",
            "extra": "134048 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3408,
            "unit": "B/op",
            "extra": "134048 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "134048 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 8940,
            "unit": "ns/op\t    5904 B/op\t      54 allocs/op",
            "extra": "131164 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 8940,
            "unit": "ns/op",
            "extra": "131164 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5904,
            "unit": "B/op",
            "extra": "131164 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "131164 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1066,
            "unit": "ns/op\t     752 B/op\t      25 allocs/op",
            "extra": "985263 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1066,
            "unit": "ns/op",
            "extra": "985263 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "985263 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "985263 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP",
            "value": 1599,
            "unit": "ns/op\t     824 B/op\t      31 allocs/op",
            "extra": "694366 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1599,
            "unit": "ns/op",
            "extra": "694366 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 824,
            "unit": "B/op",
            "extra": "694366 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "694366 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1730,
            "unit": "ns/op\t     896 B/op\t      32 allocs/op",
            "extra": "651832 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1730,
            "unit": "ns/op",
            "extra": "651832 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "651832 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "651832 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1835,
            "unit": "ns/op\t    1056 B/op\t      36 allocs/op",
            "extra": "610972 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1835,
            "unit": "ns/op",
            "extra": "610972 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "610972 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "610972 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 3034,
            "unit": "ns/op\t    1576 B/op\t      63 allocs/op",
            "extra": "373778 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 3034,
            "unit": "ns/op",
            "extra": "373778 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "373778 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "373778 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5639,
            "unit": "ns/op\t    3048 B/op\t     118 allocs/op",
            "extra": "206878 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5639,
            "unit": "ns/op",
            "extra": "206878 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 3048,
            "unit": "B/op",
            "extra": "206878 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "206878 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3566,
            "unit": "ns/op\t    2056 B/op\t      75 allocs/op",
            "extra": "322130 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3566,
            "unit": "ns/op",
            "extra": "322130 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 2056,
            "unit": "B/op",
            "extra": "322130 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "322130 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1182,
            "unit": "ns/op\t     616 B/op\t      23 allocs/op",
            "extra": "869352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1182,
            "unit": "ns/op",
            "extra": "869352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 616,
            "unit": "B/op",
            "extra": "869352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "869352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2352,
            "unit": "ns/op\t    1448 B/op\t      44 allocs/op",
            "extra": "485684 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2352,
            "unit": "ns/op",
            "extra": "485684 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "485684 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "485684 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2452,
            "unit": "ns/op\t    1576 B/op\t      46 allocs/op",
            "extra": "476518 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2452,
            "unit": "ns/op",
            "extra": "476518 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "476518 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "476518 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1313,
            "unit": "ns/op\t     784 B/op\t      24 allocs/op",
            "extra": "831379 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1313,
            "unit": "ns/op",
            "extra": "831379 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 784,
            "unit": "B/op",
            "extra": "831379 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "831379 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1308,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "838936 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1308,
            "unit": "ns/op",
            "extra": "838936 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "838936 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "838936 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1604,
            "unit": "ns/op\t     976 B/op\t      30 allocs/op",
            "extra": "704677 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1604,
            "unit": "ns/op",
            "extra": "704677 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "704677 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "704677 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 11796,
            "unit": "ns/op\t    3984 B/op\t      72 allocs/op",
            "extra": "99859 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 11796,
            "unit": "ns/op",
            "extra": "99859 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 3984,
            "unit": "B/op",
            "extra": "99859 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 72,
            "unit": "allocs/op",
            "extra": "99859 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 37348,
            "unit": "ns/op\t    7577 B/op\t      99 allocs/op",
            "extra": "32112 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 37348,
            "unit": "ns/op",
            "extra": "32112 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7577,
            "unit": "B/op",
            "extra": "32112 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "32112 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 37704,
            "unit": "ns/op\t    7833 B/op\t     108 allocs/op",
            "extra": "31657 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 37704,
            "unit": "ns/op",
            "extra": "31657 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7833,
            "unit": "B/op",
            "extra": "31657 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "31657 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 37365,
            "unit": "ns/op\t    7418 B/op\t     103 allocs/op",
            "extra": "32008 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 37365,
            "unit": "ns/op",
            "extra": "32008 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7418,
            "unit": "B/op",
            "extra": "32008 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "32008 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 16240,
            "unit": "ns/op\t    8721 B/op\t     217 allocs/op",
            "extra": "73786 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 16240,
            "unit": "ns/op",
            "extra": "73786 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8721,
            "unit": "B/op",
            "extra": "73786 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "73786 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 162610838,
            "unit": "ns/op\t334257597 B/op\t  281357 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 162610838,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334257597,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281357,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 47297,
            "unit": "ns/op\t   29417 B/op\t     614 allocs/op",
            "extra": "25264 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 47297,
            "unit": "ns/op",
            "extra": "25264 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 29417,
            "unit": "B/op",
            "extra": "25264 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "25264 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4230,
            "unit": "ns/op\t    2054 B/op\t      58 allocs/op",
            "extra": "272706 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4230,
            "unit": "ns/op",
            "extra": "272706 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 2054,
            "unit": "B/op",
            "extra": "272706 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "272706 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1315,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "845832 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1315,
            "unit": "ns/op",
            "extra": "845832 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "845832 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "845832 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "48822818+nieomylnieja@users.noreply.github.com",
            "name": "Mateusz Hawrus",
            "username": "nieomylnieja"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "733b0886f5e12e1d0399eb839f7baea8903d556d",
          "message": "chore: Add OpenSLO use case to README.md (#48)",
          "timestamp": "2024-11-07T00:14:16+01:00",
          "tree_id": "058a5c83cd3f4cf791feb3f45fb0cae9fa70ecfd",
          "url": "https://github.com/nobl9/govy/commit/733b0886f5e12e1d0399eb839f7baea8903d556d"
        },
        "date": 1730934996706,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 688.5,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1744842 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 688.5,
            "unit": "ns/op",
            "extra": "1744842 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1744842 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1744842 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 850.3,
            "unit": "ns/op\t     240 B/op\t       6 allocs/op",
            "extra": "1491212 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 850.3,
            "unit": "ns/op",
            "extra": "1491212 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "1491212 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1491212 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 850.4,
            "unit": "ns/op\t     400 B/op\t      10 allocs/op",
            "extra": "1408524 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 850.4,
            "unit": "ns/op",
            "extra": "1408524 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 400,
            "unit": "B/op",
            "extra": "1408524 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1408524 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 779.7,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1541428 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 779.7,
            "unit": "ns/op",
            "extra": "1541428 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1541428 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1541428 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 816.5,
            "unit": "ns/op\t     376 B/op\t      10 allocs/op",
            "extra": "1451911 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 816.5,
            "unit": "ns/op",
            "extra": "1451911 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 376,
            "unit": "B/op",
            "extra": "1451911 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1451911 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 787.7,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1522886 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 787.7,
            "unit": "ns/op",
            "extra": "1522886 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1522886 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1522886 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1249,
            "unit": "ns/op\t     520 B/op\t      18 allocs/op",
            "extra": "858313 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1249,
            "unit": "ns/op",
            "extra": "858313 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "858313 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "858313 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 172.1,
            "unit": "ns/op\t     144 B/op\t       4 allocs/op",
            "extra": "6728708 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 172.1,
            "unit": "ns/op",
            "extra": "6728708 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 144,
            "unit": "B/op",
            "extra": "6728708 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6728708 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1383,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "810032 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1383,
            "unit": "ns/op",
            "extra": "810032 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "810032 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "810032 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength",
            "value": 1047,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - ns/op",
            "value": 1047,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength",
            "value": 1050,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "983659 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1050,
            "unit": "ns/op",
            "extra": "983659 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "983659 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "983659 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength",
            "value": 1279,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "846560 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1279,
            "unit": "ns/op",
            "extra": "846560 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "846560 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "846560 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 1047,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 1047,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength",
            "value": 1087,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "965670 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - ns/op",
            "value": 1087,
            "unit": "ns/op",
            "extra": "965670 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "965670 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "965670 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength",
            "value": 1092,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "972364 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1092,
            "unit": "ns/op",
            "extra": "972364 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - B/op",
            "value": 560,
            "unit": "B/op",
            "extra": "972364 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "972364 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength",
            "value": 1053,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - ns/op",
            "value": 1053,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength",
            "value": 1129,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "956876 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - ns/op",
            "value": 1129,
            "unit": "ns/op",
            "extra": "956876 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "956876 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "956876 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf",
            "value": 1123,
            "unit": "ns/op\t     536 B/op\t      22 allocs/op",
            "extra": "942706 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1123,
            "unit": "ns/op",
            "extra": "942706 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 536,
            "unit": "B/op",
            "extra": "942706 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "942706 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 7369,
            "unit": "ns/op\t    3168 B/op\t      98 allocs/op",
            "extra": "160834 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 7369,
            "unit": "ns/op",
            "extra": "160834 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3168,
            "unit": "B/op",
            "extra": "160834 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "160834 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2673,
            "unit": "ns/op\t    1064 B/op\t      32 allocs/op",
            "extra": "426158 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2673,
            "unit": "ns/op",
            "extra": "426158 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1064,
            "unit": "B/op",
            "extra": "426158 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "426158 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1066,
            "unit": "ns/op\t     704 B/op\t      23 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1066,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - B/op",
            "value": 704,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty",
            "value": 181.8,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "6526372 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 181.8,
            "unit": "ns/op",
            "extra": "6526372 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "6526372 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6526372 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 1531,
            "unit": "ns/op\t     644 B/op\t      20 allocs/op",
            "extra": "732646 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 1531,
            "unit": "ns/op",
            "extra": "732646 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 644,
            "unit": "B/op",
            "extra": "732646 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "732646 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 1542,
            "unit": "ns/op\t     693 B/op\t      20 allocs/op",
            "extra": "716335 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 1542,
            "unit": "ns/op",
            "extra": "716335 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 693,
            "unit": "B/op",
            "extra": "716335 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "716335 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 15471,
            "unit": "ns/op\t    5639 B/op\t     154 allocs/op",
            "extra": "76784 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 15471,
            "unit": "ns/op",
            "extra": "76784 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 5639,
            "unit": "B/op",
            "extra": "76784 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 154,
            "unit": "allocs/op",
            "extra": "76784 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4463,
            "unit": "ns/op\t    1552 B/op\t      41 allocs/op",
            "extra": "258886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4463,
            "unit": "ns/op",
            "extra": "258886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1552,
            "unit": "B/op",
            "extra": "258886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "258886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 16384,
            "unit": "ns/op\t   14085 B/op\t     217 allocs/op",
            "extra": "72500 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 16384,
            "unit": "ns/op",
            "extra": "72500 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 14085,
            "unit": "B/op",
            "extra": "72500 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "72500 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 8778,
            "unit": "ns/op\t    3408 B/op\t     138 allocs/op",
            "extra": "134600 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 8778,
            "unit": "ns/op",
            "extra": "134600 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3408,
            "unit": "B/op",
            "extra": "134600 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "134600 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 8937,
            "unit": "ns/op\t    5904 B/op\t      54 allocs/op",
            "extra": "133954 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 8937,
            "unit": "ns/op",
            "extra": "133954 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5904,
            "unit": "B/op",
            "extra": "133954 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "133954 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1067,
            "unit": "ns/op\t     752 B/op\t      25 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1067,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP",
            "value": 1612,
            "unit": "ns/op\t     824 B/op\t      31 allocs/op",
            "extra": "688362 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1612,
            "unit": "ns/op",
            "extra": "688362 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 824,
            "unit": "B/op",
            "extra": "688362 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "688362 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1733,
            "unit": "ns/op\t     896 B/op\t      32 allocs/op",
            "extra": "657186 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1733,
            "unit": "ns/op",
            "extra": "657186 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "657186 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "657186 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1857,
            "unit": "ns/op\t    1056 B/op\t      36 allocs/op",
            "extra": "610816 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1857,
            "unit": "ns/op",
            "extra": "610816 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "610816 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "610816 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 3050,
            "unit": "ns/op\t    1576 B/op\t      63 allocs/op",
            "extra": "374109 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 3050,
            "unit": "ns/op",
            "extra": "374109 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "374109 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "374109 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5714,
            "unit": "ns/op\t    3048 B/op\t     118 allocs/op",
            "extra": "209074 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5714,
            "unit": "ns/op",
            "extra": "209074 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 3048,
            "unit": "B/op",
            "extra": "209074 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "209074 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3615,
            "unit": "ns/op\t    2056 B/op\t      75 allocs/op",
            "extra": "317557 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3615,
            "unit": "ns/op",
            "extra": "317557 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 2056,
            "unit": "B/op",
            "extra": "317557 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "317557 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1208,
            "unit": "ns/op\t     616 B/op\t      23 allocs/op",
            "extra": "911031 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1208,
            "unit": "ns/op",
            "extra": "911031 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 616,
            "unit": "B/op",
            "extra": "911031 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "911031 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2418,
            "unit": "ns/op\t    1448 B/op\t      44 allocs/op",
            "extra": "475111 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2418,
            "unit": "ns/op",
            "extra": "475111 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "475111 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "475111 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2444,
            "unit": "ns/op\t    1576 B/op\t      46 allocs/op",
            "extra": "460820 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2444,
            "unit": "ns/op",
            "extra": "460820 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "460820 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "460820 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1351,
            "unit": "ns/op\t     784 B/op\t      24 allocs/op",
            "extra": "794065 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1351,
            "unit": "ns/op",
            "extra": "794065 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 784,
            "unit": "B/op",
            "extra": "794065 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "794065 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1347,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "812775 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1347,
            "unit": "ns/op",
            "extra": "812775 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "812775 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "812775 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1625,
            "unit": "ns/op\t     976 B/op\t      30 allocs/op",
            "extra": "693416 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1625,
            "unit": "ns/op",
            "extra": "693416 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "693416 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "693416 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 11843,
            "unit": "ns/op\t    3984 B/op\t      72 allocs/op",
            "extra": "100112 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 11843,
            "unit": "ns/op",
            "extra": "100112 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 3984,
            "unit": "B/op",
            "extra": "100112 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 72,
            "unit": "allocs/op",
            "extra": "100112 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 37110,
            "unit": "ns/op\t    7529 B/op\t      99 allocs/op",
            "extra": "32113 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 37110,
            "unit": "ns/op",
            "extra": "32113 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7529,
            "unit": "B/op",
            "extra": "32113 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "32113 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 37614,
            "unit": "ns/op\t    7833 B/op\t     108 allocs/op",
            "extra": "31731 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 37614,
            "unit": "ns/op",
            "extra": "31731 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7833,
            "unit": "B/op",
            "extra": "31731 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "31731 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 37167,
            "unit": "ns/op\t    7417 B/op\t     103 allocs/op",
            "extra": "32115 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 37167,
            "unit": "ns/op",
            "extra": "32115 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7417,
            "unit": "B/op",
            "extra": "32115 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "32115 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 16481,
            "unit": "ns/op\t    8721 B/op\t     217 allocs/op",
            "extra": "71650 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 16481,
            "unit": "ns/op",
            "extra": "71650 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8721,
            "unit": "B/op",
            "extra": "71650 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "71650 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 173326670,
            "unit": "ns/op\t334255790 B/op\t  281345 allocs/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 173326670,
            "unit": "ns/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334255790,
            "unit": "B/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281345,
            "unit": "allocs/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 47772,
            "unit": "ns/op\t   29417 B/op\t     614 allocs/op",
            "extra": "24974 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 47772,
            "unit": "ns/op",
            "extra": "24974 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 29417,
            "unit": "B/op",
            "extra": "24974 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "24974 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4347,
            "unit": "ns/op\t    2054 B/op\t      58 allocs/op",
            "extra": "267912 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4347,
            "unit": "ns/op",
            "extra": "267912 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 2054,
            "unit": "B/op",
            "extra": "267912 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "267912 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1327,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "841566 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1327,
            "unit": "ns/op",
            "extra": "841566 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "841566 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "841566 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "29139614+renovate[bot]@users.noreply.github.com",
            "name": "renovate[bot]",
            "username": "renovate[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "7f9c5b3600d8439c8bff5d959b5f9b4f0eb69328",
          "message": "chore: Update dependency cspell to v8.16.0 (#49)\n\nThis PR contains the following updates:\n\n| Package | Change | Age | Adoption | Passing | Confidence |\n|---|---|---|---|---|---|\n| [cspell](https://cspell.org/)\n([source](https://redirect.github.com/streetsidesoftware/cspell/tree/HEAD/packages/cspell))\n| [`8.15.7` ->\n`8.16.0`](https://renovatebot.com/diffs/npm/cspell/8.15.7/8.16.0) |\n[![age](https://developer.mend.io/api/mc/badges/age/npm/cspell/8.16.0?slim=true)](https://docs.renovatebot.com/merge-confidence/)\n|\n[![adoption](https://developer.mend.io/api/mc/badges/adoption/npm/cspell/8.16.0?slim=true)](https://docs.renovatebot.com/merge-confidence/)\n|\n[![passing](https://developer.mend.io/api/mc/badges/compatibility/npm/cspell/8.15.7/8.16.0?slim=true)](https://docs.renovatebot.com/merge-confidence/)\n|\n[![confidence](https://developer.mend.io/api/mc/badges/confidence/npm/cspell/8.15.7/8.16.0?slim=true)](https://docs.renovatebot.com/merge-confidence/)\n|\n\n---\n\n### Release Notes\n\n<details>\n<summary>streetsidesoftware/cspell (cspell)</summary>\n\n###\n[`v8.16.0`](https://redirect.github.com/streetsidesoftware/cspell/blob/HEAD/packages/cspell/CHANGELOG.md#8160-2024-11-07)\n\n[Compare\nSource](https://redirect.github.com/streetsidesoftware/cspell/compare/v8.15.7...v8.16.0)\n\n- chore: Update Integration Test Performance Data\n([#&#8203;6505](https://redirect.github.com/streetsidesoftware/cspell/issues/6505))\n([fb78a40](https://redirect.github.com/streetsidesoftware/cspell/commit/fb78a40)),\ncloses\n[#&#8203;6505](https://redirect.github.com/streetsidesoftware/cspell/issues/6505)\n\n</details>\n\n---\n\n### Configuration\n\n📅 **Schedule**: Branch creation - \"after 10pm every weekday,before 5am\nevery weekday,every weekend\" (UTC), Automerge - At any time (no schedule\ndefined).\n\n🚦 **Automerge**: Enabled.\n\n♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the\nrebase/retry checkbox.\n\n🔕 **Ignore**: Close this PR and you won't be reminded about this update\nagain.\n\n---\n\n- [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check\nthis box\n\n---\n\nThis PR was generated by [Mend Renovate](https://mend.io/renovate/).\nView the [repository job\nlog](https://developer.mend.io/github/nobl9/govy).\n\n<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS43LjEiLCJ1cGRhdGVkSW5WZXIiOiIzOS43LjEiLCJ0YXJnZXRCcmFuY2giOiJtYWluIiwibGFiZWxzIjpbImRlcGVuZGVuY2llcyIsImphdmFzY3JpcHQiLCJyZW5vdmF0ZSJdfQ==-->\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2024-11-08T11:52:25+01:00",
          "tree_id": "b0614d3da7843144d37a6647ad67d5d1e91b321c",
          "url": "https://github.com/nobl9/govy/commit/7f9c5b3600d8439c8bff5d959b5f9b4f0eb69328"
        },
        "date": 1731063286762,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 701,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1491636 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 701,
            "unit": "ns/op",
            "extra": "1491636 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1491636 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1491636 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 813.1,
            "unit": "ns/op\t     240 B/op\t       6 allocs/op",
            "extra": "1457371 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 813.1,
            "unit": "ns/op",
            "extra": "1457371 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "1457371 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1457371 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 937.2,
            "unit": "ns/op\t     400 B/op\t      10 allocs/op",
            "extra": "1291054 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 937.2,
            "unit": "ns/op",
            "extra": "1291054 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 400,
            "unit": "B/op",
            "extra": "1291054 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1291054 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 817.9,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1496382 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 817.9,
            "unit": "ns/op",
            "extra": "1496382 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1496382 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1496382 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 936.4,
            "unit": "ns/op\t     376 B/op\t      10 allocs/op",
            "extra": "1354173 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 936.4,
            "unit": "ns/op",
            "extra": "1354173 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 376,
            "unit": "B/op",
            "extra": "1354173 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1354173 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 826.4,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1442082 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 826.4,
            "unit": "ns/op",
            "extra": "1442082 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1442082 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1442082 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1291,
            "unit": "ns/op\t     520 B/op\t      18 allocs/op",
            "extra": "908389 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1291,
            "unit": "ns/op",
            "extra": "908389 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "908389 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "908389 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 195.7,
            "unit": "ns/op\t     144 B/op\t       4 allocs/op",
            "extra": "6038287 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 195.7,
            "unit": "ns/op",
            "extra": "6038287 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 144,
            "unit": "B/op",
            "extra": "6038287 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6038287 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1465,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "836347 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1465,
            "unit": "ns/op",
            "extra": "836347 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "836347 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "836347 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength",
            "value": 1085,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1076874 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - ns/op",
            "value": 1085,
            "unit": "ns/op",
            "extra": "1076874 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "1076874 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1076874 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength",
            "value": 1070,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1070,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength",
            "value": 1313,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "848996 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1313,
            "unit": "ns/op",
            "extra": "848996 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "848996 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "848996 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 1115,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1131094 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 1115,
            "unit": "ns/op",
            "extra": "1131094 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1131094 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1131094 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength",
            "value": 1114,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "990938 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - ns/op",
            "value": 1114,
            "unit": "ns/op",
            "extra": "990938 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "990938 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "990938 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength",
            "value": 1144,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1144,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - B/op",
            "value": 560,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength",
            "value": 1128,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1128283 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - ns/op",
            "value": 1128,
            "unit": "ns/op",
            "extra": "1128283 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1128283 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1128283 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength",
            "value": 1184,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "933856 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - ns/op",
            "value": 1184,
            "unit": "ns/op",
            "extra": "933856 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "933856 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "933856 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf",
            "value": 1233,
            "unit": "ns/op\t     536 B/op\t      22 allocs/op",
            "extra": "944427 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1233,
            "unit": "ns/op",
            "extra": "944427 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 536,
            "unit": "B/op",
            "extra": "944427 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "944427 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 8240,
            "unit": "ns/op\t    3168 B/op\t      98 allocs/op",
            "extra": "150776 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 8240,
            "unit": "ns/op",
            "extra": "150776 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3168,
            "unit": "B/op",
            "extra": "150776 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "150776 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2978,
            "unit": "ns/op\t    1064 B/op\t      32 allocs/op",
            "extra": "389986 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2978,
            "unit": "ns/op",
            "extra": "389986 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1064,
            "unit": "B/op",
            "extra": "389986 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "389986 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1188,
            "unit": "ns/op\t     704 B/op\t      23 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1188,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - B/op",
            "value": 704,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty",
            "value": 196.6,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "5702179 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 196.6,
            "unit": "ns/op",
            "extra": "5702179 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "5702179 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "5702179 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 1607,
            "unit": "ns/op\t     644 B/op\t      20 allocs/op",
            "extra": "731877 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 1607,
            "unit": "ns/op",
            "extra": "731877 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 644,
            "unit": "B/op",
            "extra": "731877 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "731877 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 1584,
            "unit": "ns/op\t     693 B/op\t      20 allocs/op",
            "extra": "728938 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 1584,
            "unit": "ns/op",
            "extra": "728938 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 693,
            "unit": "B/op",
            "extra": "728938 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "728938 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 15233,
            "unit": "ns/op\t    5639 B/op\t     154 allocs/op",
            "extra": "76986 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 15233,
            "unit": "ns/op",
            "extra": "76986 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 5639,
            "unit": "B/op",
            "extra": "76986 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 154,
            "unit": "allocs/op",
            "extra": "76986 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4380,
            "unit": "ns/op\t    1552 B/op\t      41 allocs/op",
            "extra": "260977 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4380,
            "unit": "ns/op",
            "extra": "260977 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1552,
            "unit": "B/op",
            "extra": "260977 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "260977 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 17818,
            "unit": "ns/op\t   14085 B/op\t     217 allocs/op",
            "extra": "73980 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 17818,
            "unit": "ns/op",
            "extra": "73980 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 14085,
            "unit": "B/op",
            "extra": "73980 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "73980 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 9267,
            "unit": "ns/op\t    3408 B/op\t     138 allocs/op",
            "extra": "121978 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 9267,
            "unit": "ns/op",
            "extra": "121978 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3408,
            "unit": "B/op",
            "extra": "121978 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "121978 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 8982,
            "unit": "ns/op\t    5904 B/op\t      54 allocs/op",
            "extra": "128541 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 8982,
            "unit": "ns/op",
            "extra": "128541 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5904,
            "unit": "B/op",
            "extra": "128541 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "128541 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1067,
            "unit": "ns/op\t     752 B/op\t      25 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1067,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP",
            "value": 1607,
            "unit": "ns/op\t     824 B/op\t      31 allocs/op",
            "extra": "683660 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1607,
            "unit": "ns/op",
            "extra": "683660 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 824,
            "unit": "B/op",
            "extra": "683660 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "683660 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1713,
            "unit": "ns/op\t     896 B/op\t      32 allocs/op",
            "extra": "672416 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1713,
            "unit": "ns/op",
            "extra": "672416 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "672416 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "672416 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1838,
            "unit": "ns/op\t    1056 B/op\t      36 allocs/op",
            "extra": "610467 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1838,
            "unit": "ns/op",
            "extra": "610467 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "610467 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "610467 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 3016,
            "unit": "ns/op\t    1576 B/op\t      63 allocs/op",
            "extra": "374949 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 3016,
            "unit": "ns/op",
            "extra": "374949 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "374949 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "374949 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5661,
            "unit": "ns/op\t    3048 B/op\t     118 allocs/op",
            "extra": "209020 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5661,
            "unit": "ns/op",
            "extra": "209020 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 3048,
            "unit": "B/op",
            "extra": "209020 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "209020 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3553,
            "unit": "ns/op\t    2056 B/op\t      75 allocs/op",
            "extra": "311660 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3553,
            "unit": "ns/op",
            "extra": "311660 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 2056,
            "unit": "B/op",
            "extra": "311660 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "311660 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1181,
            "unit": "ns/op\t     616 B/op\t      23 allocs/op",
            "extra": "917410 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1181,
            "unit": "ns/op",
            "extra": "917410 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 616,
            "unit": "B/op",
            "extra": "917410 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "917410 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2356,
            "unit": "ns/op\t    1448 B/op\t      44 allocs/op",
            "extra": "490261 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2356,
            "unit": "ns/op",
            "extra": "490261 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "490261 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "490261 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2446,
            "unit": "ns/op\t    1576 B/op\t      46 allocs/op",
            "extra": "471645 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2446,
            "unit": "ns/op",
            "extra": "471645 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "471645 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "471645 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1356,
            "unit": "ns/op\t     784 B/op\t      24 allocs/op",
            "extra": "822439 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1356,
            "unit": "ns/op",
            "extra": "822439 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 784,
            "unit": "B/op",
            "extra": "822439 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "822439 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1323,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "807715 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1323,
            "unit": "ns/op",
            "extra": "807715 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "807715 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "807715 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1693,
            "unit": "ns/op\t     976 B/op\t      30 allocs/op",
            "extra": "684552 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1693,
            "unit": "ns/op",
            "extra": "684552 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "684552 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "684552 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 12060,
            "unit": "ns/op\t    3984 B/op\t      72 allocs/op",
            "extra": "98484 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 12060,
            "unit": "ns/op",
            "extra": "98484 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 3984,
            "unit": "B/op",
            "extra": "98484 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 72,
            "unit": "allocs/op",
            "extra": "98484 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 38905,
            "unit": "ns/op\t    7529 B/op\t      99 allocs/op",
            "extra": "30522 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 38905,
            "unit": "ns/op",
            "extra": "30522 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7529,
            "unit": "B/op",
            "extra": "30522 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "30522 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 38030,
            "unit": "ns/op\t    7833 B/op\t     108 allocs/op",
            "extra": "30742 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 38030,
            "unit": "ns/op",
            "extra": "30742 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7833,
            "unit": "B/op",
            "extra": "30742 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "30742 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 39268,
            "unit": "ns/op\t    7450 B/op\t     103 allocs/op",
            "extra": "30798 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 39268,
            "unit": "ns/op",
            "extra": "30798 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7450,
            "unit": "B/op",
            "extra": "30798 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "30798 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 17234,
            "unit": "ns/op\t    8721 B/op\t     217 allocs/op",
            "extra": "68530 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 17234,
            "unit": "ns/op",
            "extra": "68530 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8721,
            "unit": "B/op",
            "extra": "68530 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "68530 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 181423526,
            "unit": "ns/op\t334258436 B/op\t  281363 allocs/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 181423526,
            "unit": "ns/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334258436,
            "unit": "B/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281363,
            "unit": "allocs/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 52109,
            "unit": "ns/op\t   29417 B/op\t     614 allocs/op",
            "extra": "22954 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 52109,
            "unit": "ns/op",
            "extra": "22954 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 29417,
            "unit": "B/op",
            "extra": "22954 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "22954 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4491,
            "unit": "ns/op\t    2054 B/op\t      58 allocs/op",
            "extra": "248472 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4491,
            "unit": "ns/op",
            "extra": "248472 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 2054,
            "unit": "B/op",
            "extra": "248472 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "248472 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1343,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "905860 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1343,
            "unit": "ns/op",
            "extra": "905860 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "905860 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "905860 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "48822818+nieomylnieja@users.noreply.github.com",
            "name": "Mateusz Hawrus",
            "username": "nieomylnieja"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "2b70de379c99c141b9866385ea78eef5be5914ec",
          "message": "feat: Add StringDateTime and StringTimeZone rules (#50)\n\n## Release Notes\r\n\r\nAdded `StringDateTime` rule which ensures a string is a valid date and\r\ntime according to the rules defined by https://pkg.go.dev/time.\r\nAdded `StringTimeZone` rule which ensures a string is a valid IANA Time\r\nZone database code.",
          "timestamp": "2024-11-09T17:44:46+01:00",
          "tree_id": "00b70eb367d39736f8eced6fac3f12d5f2f8f321",
          "url": "https://github.com/nobl9/govy/commit/2b70de379c99c141b9866385ea78eef5be5914ec"
        },
        "date": 1731170834054,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 672.1,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1786438 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 672.1,
            "unit": "ns/op",
            "extra": "1786438 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1786438 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1786438 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 762.1,
            "unit": "ns/op\t     240 B/op\t       6 allocs/op",
            "extra": "1586227 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 762.1,
            "unit": "ns/op",
            "extra": "1586227 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "1586227 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1586227 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 810.2,
            "unit": "ns/op\t     400 B/op\t      10 allocs/op",
            "extra": "1480088 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 810.2,
            "unit": "ns/op",
            "extra": "1480088 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 400,
            "unit": "B/op",
            "extra": "1480088 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1480088 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 760.9,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1575350 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 760.9,
            "unit": "ns/op",
            "extra": "1575350 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1575350 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1575350 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 807.6,
            "unit": "ns/op\t     376 B/op\t      10 allocs/op",
            "extra": "1364274 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 807.6,
            "unit": "ns/op",
            "extra": "1364274 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 376,
            "unit": "B/op",
            "extra": "1364274 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1364274 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 750,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1598167 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 750,
            "unit": "ns/op",
            "extra": "1598167 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1598167 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1598167 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1191,
            "unit": "ns/op\t     520 B/op\t      18 allocs/op",
            "extra": "885140 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1191,
            "unit": "ns/op",
            "extra": "885140 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "885140 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "885140 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 173.2,
            "unit": "ns/op\t     144 B/op\t       4 allocs/op",
            "extra": "7003464 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 173.2,
            "unit": "ns/op",
            "extra": "7003464 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 144,
            "unit": "B/op",
            "extra": "7003464 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "7003464 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1249,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "863352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1249,
            "unit": "ns/op",
            "extra": "863352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "863352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "863352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength",
            "value": 1064,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "956494 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - ns/op",
            "value": 1064,
            "unit": "ns/op",
            "extra": "956494 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "956494 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "956494 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength",
            "value": 1045,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1045,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength",
            "value": 1244,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "879422 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1244,
            "unit": "ns/op",
            "extra": "879422 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "879422 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "879422 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 1008,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 1008,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength",
            "value": 1008,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - ns/op",
            "value": 1008,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength",
            "value": 1053,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1053,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - B/op",
            "value": 560,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength",
            "value": 1010,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - ns/op",
            "value": 1010,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength",
            "value": 1006,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - ns/op",
            "value": 1006,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf",
            "value": 1114,
            "unit": "ns/op\t     536 B/op\t      22 allocs/op",
            "extra": "978493 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1114,
            "unit": "ns/op",
            "extra": "978493 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 536,
            "unit": "B/op",
            "extra": "978493 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "978493 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 6977,
            "unit": "ns/op\t    3168 B/op\t      98 allocs/op",
            "extra": "167682 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 6977,
            "unit": "ns/op",
            "extra": "167682 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3168,
            "unit": "B/op",
            "extra": "167682 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "167682 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2500,
            "unit": "ns/op\t    1064 B/op\t      32 allocs/op",
            "extra": "454808 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2500,
            "unit": "ns/op",
            "extra": "454808 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1064,
            "unit": "B/op",
            "extra": "454808 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "454808 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1084,
            "unit": "ns/op\t     704 B/op\t      23 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1084,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - B/op",
            "value": 704,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty",
            "value": 179.3,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "6604903 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 179.3,
            "unit": "ns/op",
            "extra": "6604903 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "6604903 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6604903 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 1523,
            "unit": "ns/op\t     644 B/op\t      20 allocs/op",
            "extra": "714368 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 1523,
            "unit": "ns/op",
            "extra": "714368 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 644,
            "unit": "B/op",
            "extra": "714368 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "714368 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 1519,
            "unit": "ns/op\t     692 B/op\t      20 allocs/op",
            "extra": "742620 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 1519,
            "unit": "ns/op",
            "extra": "742620 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 692,
            "unit": "B/op",
            "extra": "742620 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "742620 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 15680,
            "unit": "ns/op\t    5643 B/op\t     154 allocs/op",
            "extra": "76148 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 15680,
            "unit": "ns/op",
            "extra": "76148 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 5643,
            "unit": "B/op",
            "extra": "76148 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 154,
            "unit": "allocs/op",
            "extra": "76148 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4376,
            "unit": "ns/op\t    1552 B/op\t      41 allocs/op",
            "extra": "264535 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4376,
            "unit": "ns/op",
            "extra": "264535 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1552,
            "unit": "B/op",
            "extra": "264535 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "264535 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 16506,
            "unit": "ns/op\t   14085 B/op\t     217 allocs/op",
            "extra": "72278 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 16506,
            "unit": "ns/op",
            "extra": "72278 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 14085,
            "unit": "B/op",
            "extra": "72278 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "72278 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 8629,
            "unit": "ns/op\t    3408 B/op\t     138 allocs/op",
            "extra": "135200 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 8629,
            "unit": "ns/op",
            "extra": "135200 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3408,
            "unit": "B/op",
            "extra": "135200 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "135200 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 9229,
            "unit": "ns/op\t    5904 B/op\t      54 allocs/op",
            "extra": "127461 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 9229,
            "unit": "ns/op",
            "extra": "127461 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5904,
            "unit": "B/op",
            "extra": "127461 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "127461 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1066,
            "unit": "ns/op\t     752 B/op\t      25 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1066,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP",
            "value": 1594,
            "unit": "ns/op\t     824 B/op\t      31 allocs/op",
            "extra": "688111 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1594,
            "unit": "ns/op",
            "extra": "688111 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 824,
            "unit": "B/op",
            "extra": "688111 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "688111 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1683,
            "unit": "ns/op\t     896 B/op\t      32 allocs/op",
            "extra": "659595 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1683,
            "unit": "ns/op",
            "extra": "659595 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "659595 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "659595 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1863,
            "unit": "ns/op\t    1056 B/op\t      36 allocs/op",
            "extra": "601999 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1863,
            "unit": "ns/op",
            "extra": "601999 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "601999 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "601999 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 3081,
            "unit": "ns/op\t    1576 B/op\t      63 allocs/op",
            "extra": "375403 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 3081,
            "unit": "ns/op",
            "extra": "375403 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "375403 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "375403 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5637,
            "unit": "ns/op\t    3048 B/op\t     118 allocs/op",
            "extra": "208848 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5637,
            "unit": "ns/op",
            "extra": "208848 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 3048,
            "unit": "B/op",
            "extra": "208848 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "208848 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3573,
            "unit": "ns/op\t    2056 B/op\t      75 allocs/op",
            "extra": "322375 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3573,
            "unit": "ns/op",
            "extra": "322375 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 2056,
            "unit": "B/op",
            "extra": "322375 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "322375 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1193,
            "unit": "ns/op\t     616 B/op\t      23 allocs/op",
            "extra": "924542 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1193,
            "unit": "ns/op",
            "extra": "924542 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 616,
            "unit": "B/op",
            "extra": "924542 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "924542 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2360,
            "unit": "ns/op\t    1448 B/op\t      44 allocs/op",
            "extra": "482103 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2360,
            "unit": "ns/op",
            "extra": "482103 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "482103 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "482103 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2444,
            "unit": "ns/op\t    1576 B/op\t      46 allocs/op",
            "extra": "475317 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2444,
            "unit": "ns/op",
            "extra": "475317 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "475317 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "475317 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1344,
            "unit": "ns/op\t     784 B/op\t      24 allocs/op",
            "extra": "812547 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1344,
            "unit": "ns/op",
            "extra": "812547 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 784,
            "unit": "B/op",
            "extra": "812547 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "812547 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1338,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "808186 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1338,
            "unit": "ns/op",
            "extra": "808186 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "808186 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "808186 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1627,
            "unit": "ns/op\t     976 B/op\t      30 allocs/op",
            "extra": "697172 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1627,
            "unit": "ns/op",
            "extra": "697172 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "697172 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "697172 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 11742,
            "unit": "ns/op\t    3984 B/op\t      72 allocs/op",
            "extra": "101623 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 11742,
            "unit": "ns/op",
            "extra": "101623 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 3984,
            "unit": "B/op",
            "extra": "101623 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 72,
            "unit": "allocs/op",
            "extra": "101623 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 37207,
            "unit": "ns/op\t    7577 B/op\t      99 allocs/op",
            "extra": "32256 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 37207,
            "unit": "ns/op",
            "extra": "32256 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7577,
            "unit": "B/op",
            "extra": "32256 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "32256 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 37585,
            "unit": "ns/op\t    7754 B/op\t     108 allocs/op",
            "extra": "31648 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 37585,
            "unit": "ns/op",
            "extra": "31648 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7754,
            "unit": "B/op",
            "extra": "31648 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "31648 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 37466,
            "unit": "ns/op\t    7449 B/op\t     103 allocs/op",
            "extra": "32096 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 37466,
            "unit": "ns/op",
            "extra": "32096 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7449,
            "unit": "B/op",
            "extra": "32096 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "32096 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 15729,
            "unit": "ns/op\t    8721 B/op\t     217 allocs/op",
            "extra": "75258 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 15729,
            "unit": "ns/op",
            "extra": "75258 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8721,
            "unit": "B/op",
            "extra": "75258 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "75258 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 161509163,
            "unit": "ns/op\t334257302 B/op\t  281355 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 161509163,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334257302,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281355,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 47586,
            "unit": "ns/op\t   29417 B/op\t     614 allocs/op",
            "extra": "25138 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 47586,
            "unit": "ns/op",
            "extra": "25138 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 29417,
            "unit": "B/op",
            "extra": "25138 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "25138 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime",
            "value": 6214,
            "unit": "ns/op\t    3104 B/op\t      76 allocs/op",
            "extra": "187380 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - ns/op",
            "value": 6214,
            "unit": "ns/op",
            "extra": "187380 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - B/op",
            "value": 3104,
            "unit": "B/op",
            "extra": "187380 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - allocs/op",
            "value": 76,
            "unit": "allocs/op",
            "extra": "187380 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone",
            "value": 228967,
            "unit": "ns/op\t  337430 B/op\t     224 allocs/op",
            "extra": "4887 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - ns/op",
            "value": 228967,
            "unit": "ns/op",
            "extra": "4887 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - B/op",
            "value": 337430,
            "unit": "B/op",
            "extra": "4887 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "4887 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4194,
            "unit": "ns/op\t    2054 B/op\t      58 allocs/op",
            "extra": "274576 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4194,
            "unit": "ns/op",
            "extra": "274576 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 2054,
            "unit": "B/op",
            "extra": "274576 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "274576 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1307,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "843986 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1307,
            "unit": "ns/op",
            "extra": "843986 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "843986 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "843986 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "48822818+nieomylnieja@users.noreply.github.com",
            "name": "Mateusz Hawrus",
            "username": "nieomylnieja"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "070e14fc62d3492684500c05a2147780c01d3525",
          "message": "fix: Predicate matching for govy.PropertyRules (#51)\n\n## Summary\r\n\r\n`PropertyRulesForMap` and `PropertyRulesForSlice` already had the\r\npredicates matching done before a value getter was called.\r\nHaving this check done for `PropertyRules` **AFTER** extracting the\r\nvalue was causing `Required` to fail with an error before any conditions\r\nwere checked.\r\n\r\nExample of such scenario:\r\n\r\n```go\r\nr := govy.ForPointer(func(s *string) *string { return s }).\r\n\tWhen(func(s *string) bool { return s != nil }).\r\n\tRequired().\r\n\tRules(rules.StringMinLength(10))\r\nerr := r.Validate(nil)\r\nassert.NoError(t, err) // Fails!\r\n```\r\n\r\n## Breaking Changes\r\n\r\n`govy.PropertyRules` will no longer fail if `Required()` was specified\r\nwhen a validate value is its type's zero value **AND** none of the\r\n`When()` conditions are matched.",
          "timestamp": "2024-11-18T12:07:02+01:00",
          "tree_id": "fcac5e165230977f2a4d5f1ceefe20261bb9abcd",
          "url": "https://github.com/nobl9/govy/commit/070e14fc62d3492684500c05a2147780c01d3525"
        },
        "date": 1731928166714,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 694.9,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1742977 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 694.9,
            "unit": "ns/op",
            "extra": "1742977 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1742977 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1742977 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 766.1,
            "unit": "ns/op\t     240 B/op\t       6 allocs/op",
            "extra": "1555413 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 766.1,
            "unit": "ns/op",
            "extra": "1555413 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "1555413 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1555413 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 798.2,
            "unit": "ns/op\t     400 B/op\t      10 allocs/op",
            "extra": "1507011 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 798.2,
            "unit": "ns/op",
            "extra": "1507011 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 400,
            "unit": "B/op",
            "extra": "1507011 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1507011 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 732.1,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1646032 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 732.1,
            "unit": "ns/op",
            "extra": "1646032 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1646032 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1646032 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 838.2,
            "unit": "ns/op\t     376 B/op\t      10 allocs/op",
            "extra": "1518270 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 838.2,
            "unit": "ns/op",
            "extra": "1518270 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 376,
            "unit": "B/op",
            "extra": "1518270 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1518270 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 722.2,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1654222 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 722.2,
            "unit": "ns/op",
            "extra": "1654222 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1654222 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1654222 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1238,
            "unit": "ns/op\t     520 B/op\t      18 allocs/op",
            "extra": "883077 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1238,
            "unit": "ns/op",
            "extra": "883077 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "883077 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "883077 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 174.2,
            "unit": "ns/op\t     144 B/op\t       4 allocs/op",
            "extra": "6924110 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 174.2,
            "unit": "ns/op",
            "extra": "6924110 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 144,
            "unit": "B/op",
            "extra": "6924110 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6924110 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1256,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "869062 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1256,
            "unit": "ns/op",
            "extra": "869062 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "869062 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "869062 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength",
            "value": 1062,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "984505 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - ns/op",
            "value": 1062,
            "unit": "ns/op",
            "extra": "984505 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "984505 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "984505 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength",
            "value": 1077,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "986966 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1077,
            "unit": "ns/op",
            "extra": "986966 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "986966 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "986966 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength",
            "value": 1261,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "880654 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1261,
            "unit": "ns/op",
            "extra": "880654 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "880654 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "880654 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 997.8,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1203468 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 997.8,
            "unit": "ns/op",
            "extra": "1203468 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1203468 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1203468 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength",
            "value": 1007,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - ns/op",
            "value": 1007,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength",
            "value": 1053,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1053,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - B/op",
            "value": 560,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength",
            "value": 1020,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - ns/op",
            "value": 1020,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength",
            "value": 1024,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - ns/op",
            "value": 1024,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf",
            "value": 1126,
            "unit": "ns/op\t     536 B/op\t      22 allocs/op",
            "extra": "959239 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1126,
            "unit": "ns/op",
            "extra": "959239 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 536,
            "unit": "B/op",
            "extra": "959239 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "959239 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 7264,
            "unit": "ns/op\t    3168 B/op\t      98 allocs/op",
            "extra": "163642 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 7264,
            "unit": "ns/op",
            "extra": "163642 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3168,
            "unit": "B/op",
            "extra": "163642 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "163642 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2604,
            "unit": "ns/op\t    1064 B/op\t      32 allocs/op",
            "extra": "444777 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2604,
            "unit": "ns/op",
            "extra": "444777 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1064,
            "unit": "B/op",
            "extra": "444777 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "444777 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1116,
            "unit": "ns/op\t     704 B/op\t      23 allocs/op",
            "extra": "961233 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1116,
            "unit": "ns/op",
            "extra": "961233 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - B/op",
            "value": 704,
            "unit": "B/op",
            "extra": "961233 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "961233 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty",
            "value": 180.5,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "6596421 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 180.5,
            "unit": "ns/op",
            "extra": "6596421 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "6596421 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6596421 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 1512,
            "unit": "ns/op\t     644 B/op\t      20 allocs/op",
            "extra": "722014 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 1512,
            "unit": "ns/op",
            "extra": "722014 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 644,
            "unit": "B/op",
            "extra": "722014 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "722014 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 1527,
            "unit": "ns/op\t     693 B/op\t      20 allocs/op",
            "extra": "726914 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 1527,
            "unit": "ns/op",
            "extra": "726914 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 693,
            "unit": "B/op",
            "extra": "726914 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "726914 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 15445,
            "unit": "ns/op\t    5638 B/op\t     154 allocs/op",
            "extra": "76582 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 15445,
            "unit": "ns/op",
            "extra": "76582 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 5638,
            "unit": "B/op",
            "extra": "76582 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 154,
            "unit": "allocs/op",
            "extra": "76582 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4391,
            "unit": "ns/op\t    1552 B/op\t      41 allocs/op",
            "extra": "264171 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4391,
            "unit": "ns/op",
            "extra": "264171 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1552,
            "unit": "B/op",
            "extra": "264171 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "264171 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 16136,
            "unit": "ns/op\t   14085 B/op\t     217 allocs/op",
            "extra": "73578 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 16136,
            "unit": "ns/op",
            "extra": "73578 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 14085,
            "unit": "B/op",
            "extra": "73578 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "73578 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 8618,
            "unit": "ns/op\t    3408 B/op\t     138 allocs/op",
            "extra": "136030 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 8618,
            "unit": "ns/op",
            "extra": "136030 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3408,
            "unit": "B/op",
            "extra": "136030 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "136030 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 9059,
            "unit": "ns/op\t    5904 B/op\t      54 allocs/op",
            "extra": "133146 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 9059,
            "unit": "ns/op",
            "extra": "133146 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5904,
            "unit": "B/op",
            "extra": "133146 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "133146 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1088,
            "unit": "ns/op\t     752 B/op\t      25 allocs/op",
            "extra": "964935 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1088,
            "unit": "ns/op",
            "extra": "964935 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "964935 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "964935 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP",
            "value": 1570,
            "unit": "ns/op\t     824 B/op\t      31 allocs/op",
            "extra": "713335 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1570,
            "unit": "ns/op",
            "extra": "713335 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 824,
            "unit": "B/op",
            "extra": "713335 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "713335 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1692,
            "unit": "ns/op\t     896 B/op\t      32 allocs/op",
            "extra": "661158 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1692,
            "unit": "ns/op",
            "extra": "661158 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "661158 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "661158 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1837,
            "unit": "ns/op\t    1056 B/op\t      36 allocs/op",
            "extra": "614962 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1837,
            "unit": "ns/op",
            "extra": "614962 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "614962 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "614962 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 3004,
            "unit": "ns/op\t    1576 B/op\t      63 allocs/op",
            "extra": "380592 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 3004,
            "unit": "ns/op",
            "extra": "380592 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "380592 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "380592 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5544,
            "unit": "ns/op\t    3048 B/op\t     118 allocs/op",
            "extra": "210627 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5544,
            "unit": "ns/op",
            "extra": "210627 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 3048,
            "unit": "B/op",
            "extra": "210627 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "210627 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3692,
            "unit": "ns/op\t    2056 B/op\t      75 allocs/op",
            "extra": "324124 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3692,
            "unit": "ns/op",
            "extra": "324124 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 2056,
            "unit": "B/op",
            "extra": "324124 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "324124 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1156,
            "unit": "ns/op\t     616 B/op\t      23 allocs/op",
            "extra": "943518 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1156,
            "unit": "ns/op",
            "extra": "943518 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 616,
            "unit": "B/op",
            "extra": "943518 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "943518 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2373,
            "unit": "ns/op\t    1448 B/op\t      44 allocs/op",
            "extra": "474724 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2373,
            "unit": "ns/op",
            "extra": "474724 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "474724 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "474724 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2458,
            "unit": "ns/op\t    1576 B/op\t      46 allocs/op",
            "extra": "468385 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2458,
            "unit": "ns/op",
            "extra": "468385 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "468385 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "468385 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1358,
            "unit": "ns/op\t     784 B/op\t      24 allocs/op",
            "extra": "802262 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1358,
            "unit": "ns/op",
            "extra": "802262 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 784,
            "unit": "B/op",
            "extra": "802262 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "802262 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1366,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "799676 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1366,
            "unit": "ns/op",
            "extra": "799676 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "799676 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "799676 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1631,
            "unit": "ns/op\t     976 B/op\t      30 allocs/op",
            "extra": "676640 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1631,
            "unit": "ns/op",
            "extra": "676640 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "676640 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "676640 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 11774,
            "unit": "ns/op\t    3984 B/op\t      72 allocs/op",
            "extra": "101336 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 11774,
            "unit": "ns/op",
            "extra": "101336 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 3984,
            "unit": "B/op",
            "extra": "101336 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 72,
            "unit": "allocs/op",
            "extra": "101336 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 36872,
            "unit": "ns/op\t    7529 B/op\t      99 allocs/op",
            "extra": "32424 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 36872,
            "unit": "ns/op",
            "extra": "32424 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7529,
            "unit": "B/op",
            "extra": "32424 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "32424 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 37489,
            "unit": "ns/op\t    7834 B/op\t     108 allocs/op",
            "extra": "31032 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 37489,
            "unit": "ns/op",
            "extra": "31032 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7834,
            "unit": "B/op",
            "extra": "31032 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "31032 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 37122,
            "unit": "ns/op\t    7418 B/op\t     103 allocs/op",
            "extra": "32181 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 37122,
            "unit": "ns/op",
            "extra": "32181 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7418,
            "unit": "B/op",
            "extra": "32181 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "32181 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 16042,
            "unit": "ns/op\t    8721 B/op\t     217 allocs/op",
            "extra": "76222 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 16042,
            "unit": "ns/op",
            "extra": "76222 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8721,
            "unit": "B/op",
            "extra": "76222 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "76222 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 159684377,
            "unit": "ns/op\t334254265 B/op\t  281334 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 159684377,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334254265,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281334,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 46803,
            "unit": "ns/op\t   29417 B/op\t     614 allocs/op",
            "extra": "25429 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 46803,
            "unit": "ns/op",
            "extra": "25429 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 29417,
            "unit": "B/op",
            "extra": "25429 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "25429 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime",
            "value": 6330,
            "unit": "ns/op\t    3104 B/op\t      76 allocs/op",
            "extra": "189147 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - ns/op",
            "value": 6330,
            "unit": "ns/op",
            "extra": "189147 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - B/op",
            "value": 3104,
            "unit": "B/op",
            "extra": "189147 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - allocs/op",
            "value": 76,
            "unit": "allocs/op",
            "extra": "189147 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone",
            "value": 227558,
            "unit": "ns/op\t  337432 B/op\t     224 allocs/op",
            "extra": "4898 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - ns/op",
            "value": 227558,
            "unit": "ns/op",
            "extra": "4898 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - B/op",
            "value": 337432,
            "unit": "B/op",
            "extra": "4898 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "4898 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4095,
            "unit": "ns/op\t    2054 B/op\t      58 allocs/op",
            "extra": "282597 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4095,
            "unit": "ns/op",
            "extra": "282597 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 2054,
            "unit": "B/op",
            "extra": "282597 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "282597 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1309,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "855252 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1309,
            "unit": "ns/op",
            "extra": "855252 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "855252 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "855252 times\n4 procs"
          }
        ]
      }
    ]
  }
}