window.BENCHMARK_DATA = {
  "lastUpdate": 1736769200032,
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
          "id": "85c780936fec2f01051db9eb4853c9e7ad507f5f",
          "message": "feat: Add alphanumeric rules (#52)\n\n## Release Notes\r\n\r\nAdded `rules.StringAlpha`, `rules.StringAlphanumeric`,\r\n`rules.StringAlphaUnicode` and `rules.StringAlphanumericUnicode` which\r\nhelp ensure strings contain only letters and numbers (either ASCII or\r\nUnicode).",
          "timestamp": "2024-11-19T16:36:03+01:00",
          "tree_id": "52ce7bc24d160dce828a3b7d29ba405259b4f23e",
          "url": "https://github.com/nobl9/govy/commit/85c780936fec2f01051db9eb4853c9e7ad507f5f"
        },
        "date": 1732030704990,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 744.4,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1518987 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 744.4,
            "unit": "ns/op",
            "extra": "1518987 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1518987 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1518987 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 778.3,
            "unit": "ns/op\t     240 B/op\t       6 allocs/op",
            "extra": "1557394 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 778.3,
            "unit": "ns/op",
            "extra": "1557394 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "1557394 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1557394 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 819.5,
            "unit": "ns/op\t     400 B/op\t      10 allocs/op",
            "extra": "1464006 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 819.5,
            "unit": "ns/op",
            "extra": "1464006 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 400,
            "unit": "B/op",
            "extra": "1464006 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1464006 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 738.9,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1621874 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 738.9,
            "unit": "ns/op",
            "extra": "1621874 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1621874 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1621874 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 820.5,
            "unit": "ns/op\t     376 B/op\t      10 allocs/op",
            "extra": "1467460 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 820.5,
            "unit": "ns/op",
            "extra": "1467460 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 376,
            "unit": "B/op",
            "extra": "1467460 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1467460 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 750.3,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1593951 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 750.3,
            "unit": "ns/op",
            "extra": "1593951 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1593951 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1593951 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1247,
            "unit": "ns/op\t     520 B/op\t      18 allocs/op",
            "extra": "877156 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1247,
            "unit": "ns/op",
            "extra": "877156 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "877156 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "877156 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 179.7,
            "unit": "ns/op\t     144 B/op\t       4 allocs/op",
            "extra": "6858004 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 179.7,
            "unit": "ns/op",
            "extra": "6858004 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 144,
            "unit": "B/op",
            "extra": "6858004 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6858004 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1258,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "876326 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1258,
            "unit": "ns/op",
            "extra": "876326 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "876326 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "876326 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength",
            "value": 1025,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - ns/op",
            "value": 1025,
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
            "value": 1043,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "968761 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1043,
            "unit": "ns/op",
            "extra": "968761 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "968761 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "968761 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength",
            "value": 1256,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "866626 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1256,
            "unit": "ns/op",
            "extra": "866626 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "866626 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "866626 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 1020,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 1020,
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
            "value": 1012,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - ns/op",
            "value": 1012,
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
            "value": 1051,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1051,
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
            "extra": "1189214 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - ns/op",
            "value": 1010,
            "unit": "ns/op",
            "extra": "1189214 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1189214 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1189214 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength",
            "value": 1007,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - ns/op",
            "value": 1007,
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
            "value": 1133,
            "unit": "ns/op\t     536 B/op\t      22 allocs/op",
            "extra": "957279 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1133,
            "unit": "ns/op",
            "extra": "957279 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 536,
            "unit": "B/op",
            "extra": "957279 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "957279 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 7425,
            "unit": "ns/op\t    3168 B/op\t      98 allocs/op",
            "extra": "161188 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 7425,
            "unit": "ns/op",
            "extra": "161188 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3168,
            "unit": "B/op",
            "extra": "161188 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "161188 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2554,
            "unit": "ns/op\t    1064 B/op\t      32 allocs/op",
            "extra": "444798 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2554,
            "unit": "ns/op",
            "extra": "444798 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1064,
            "unit": "B/op",
            "extra": "444798 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "444798 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1087,
            "unit": "ns/op\t     704 B/op\t      23 allocs/op",
            "extra": "993517 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1087,
            "unit": "ns/op",
            "extra": "993517 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - B/op",
            "value": 704,
            "unit": "B/op",
            "extra": "993517 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "993517 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty",
            "value": 180.4,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "6585825 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 180.4,
            "unit": "ns/op",
            "extra": "6585825 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "6585825 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6585825 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 1515,
            "unit": "ns/op\t     644 B/op\t      20 allocs/op",
            "extra": "748084 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 1515,
            "unit": "ns/op",
            "extra": "748084 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 644,
            "unit": "B/op",
            "extra": "748084 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "748084 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 1562,
            "unit": "ns/op\t     693 B/op\t      20 allocs/op",
            "extra": "713462 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 1562,
            "unit": "ns/op",
            "extra": "713462 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 693,
            "unit": "B/op",
            "extra": "713462 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "713462 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 15513,
            "unit": "ns/op\t    5638 B/op\t     154 allocs/op",
            "extra": "76770 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 15513,
            "unit": "ns/op",
            "extra": "76770 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 5638,
            "unit": "B/op",
            "extra": "76770 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 154,
            "unit": "allocs/op",
            "extra": "76770 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4353,
            "unit": "ns/op\t    1552 B/op\t      41 allocs/op",
            "extra": "266290 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4353,
            "unit": "ns/op",
            "extra": "266290 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1552,
            "unit": "B/op",
            "extra": "266290 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "266290 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 15974,
            "unit": "ns/op\t   14085 B/op\t     217 allocs/op",
            "extra": "74850 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 15974,
            "unit": "ns/op",
            "extra": "74850 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 14085,
            "unit": "B/op",
            "extra": "74850 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "74850 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 8582,
            "unit": "ns/op\t    3408 B/op\t     138 allocs/op",
            "extra": "137524 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 8582,
            "unit": "ns/op",
            "extra": "137524 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3408,
            "unit": "B/op",
            "extra": "137524 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "137524 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 8936,
            "unit": "ns/op\t    5904 B/op\t      54 allocs/op",
            "extra": "133302 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 8936,
            "unit": "ns/op",
            "extra": "133302 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5904,
            "unit": "B/op",
            "extra": "133302 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "133302 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1077,
            "unit": "ns/op\t     752 B/op\t      25 allocs/op",
            "extra": "972574 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1077,
            "unit": "ns/op",
            "extra": "972574 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "972574 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "972574 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP",
            "value": 1578,
            "unit": "ns/op\t     824 B/op\t      31 allocs/op",
            "extra": "692095 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1578,
            "unit": "ns/op",
            "extra": "692095 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 824,
            "unit": "B/op",
            "extra": "692095 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "692095 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1682,
            "unit": "ns/op\t     896 B/op\t      32 allocs/op",
            "extra": "662787 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1682,
            "unit": "ns/op",
            "extra": "662787 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "662787 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "662787 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1858,
            "unit": "ns/op\t    1056 B/op\t      36 allocs/op",
            "extra": "611503 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1858,
            "unit": "ns/op",
            "extra": "611503 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "611503 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "611503 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 3026,
            "unit": "ns/op\t    1576 B/op\t      63 allocs/op",
            "extra": "378333 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 3026,
            "unit": "ns/op",
            "extra": "378333 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "378333 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "378333 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5573,
            "unit": "ns/op\t    3048 B/op\t     118 allocs/op",
            "extra": "210624 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5573,
            "unit": "ns/op",
            "extra": "210624 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 3048,
            "unit": "B/op",
            "extra": "210624 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "210624 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3596,
            "unit": "ns/op\t    2056 B/op\t      75 allocs/op",
            "extra": "323228 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3596,
            "unit": "ns/op",
            "extra": "323228 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 2056,
            "unit": "B/op",
            "extra": "323228 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "323228 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1167,
            "unit": "ns/op\t     616 B/op\t      23 allocs/op",
            "extra": "924703 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1167,
            "unit": "ns/op",
            "extra": "924703 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 616,
            "unit": "B/op",
            "extra": "924703 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "924703 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2377,
            "unit": "ns/op\t    1448 B/op\t      44 allocs/op",
            "extra": "475142 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2377,
            "unit": "ns/op",
            "extra": "475142 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "475142 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "475142 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2431,
            "unit": "ns/op\t    1576 B/op\t      46 allocs/op",
            "extra": "471042 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2431,
            "unit": "ns/op",
            "extra": "471042 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "471042 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "471042 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1355,
            "unit": "ns/op\t     784 B/op\t      24 allocs/op",
            "extra": "783528 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1355,
            "unit": "ns/op",
            "extra": "783528 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 784,
            "unit": "B/op",
            "extra": "783528 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "783528 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1343,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "799579 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1343,
            "unit": "ns/op",
            "extra": "799579 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "799579 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "799579 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1647,
            "unit": "ns/op\t     976 B/op\t      30 allocs/op",
            "extra": "684984 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1647,
            "unit": "ns/op",
            "extra": "684984 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "684984 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "684984 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 11772,
            "unit": "ns/op\t    3984 B/op\t      72 allocs/op",
            "extra": "100370 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 11772,
            "unit": "ns/op",
            "extra": "100370 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 3984,
            "unit": "B/op",
            "extra": "100370 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 72,
            "unit": "allocs/op",
            "extra": "100370 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 37018,
            "unit": "ns/op\t    7529 B/op\t      99 allocs/op",
            "extra": "32287 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 37018,
            "unit": "ns/op",
            "extra": "32287 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7529,
            "unit": "B/op",
            "extra": "32287 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "32287 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 38368,
            "unit": "ns/op\t    7753 B/op\t     108 allocs/op",
            "extra": "31742 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 38368,
            "unit": "ns/op",
            "extra": "31742 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7753,
            "unit": "B/op",
            "extra": "31742 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "31742 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 37228,
            "unit": "ns/op\t    7450 B/op\t     103 allocs/op",
            "extra": "29142 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 37228,
            "unit": "ns/op",
            "extra": "29142 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7450,
            "unit": "B/op",
            "extra": "29142 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "29142 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 16104,
            "unit": "ns/op\t    8721 B/op\t     217 allocs/op",
            "extra": "74352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 16104,
            "unit": "ns/op",
            "extra": "74352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8721,
            "unit": "B/op",
            "extra": "74352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "74352 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 166446914,
            "unit": "ns/op\t334259010 B/op\t  281367 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 166446914,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334259010,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281367,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 46699,
            "unit": "ns/op\t   29417 B/op\t     614 allocs/op",
            "extra": "25346 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 46699,
            "unit": "ns/op",
            "extra": "25346 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 29417,
            "unit": "B/op",
            "extra": "25346 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "25346 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime",
            "value": 6216,
            "unit": "ns/op\t    3104 B/op\t      76 allocs/op",
            "extra": "185251 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - ns/op",
            "value": 6216,
            "unit": "ns/op",
            "extra": "185251 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - B/op",
            "value": 3104,
            "unit": "B/op",
            "extra": "185251 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - allocs/op",
            "value": 76,
            "unit": "allocs/op",
            "extra": "185251 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone",
            "value": 231369,
            "unit": "ns/op\t  337428 B/op\t     224 allocs/op",
            "extra": "4574 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - ns/op",
            "value": 231369,
            "unit": "ns/op",
            "extra": "4574 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - B/op",
            "value": 337428,
            "unit": "B/op",
            "extra": "4574 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "4574 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha",
            "value": 3693,
            "unit": "ns/op\t    1600 B/op\t      42 allocs/op",
            "extra": "306914 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - ns/op",
            "value": 3693,
            "unit": "ns/op",
            "extra": "306914 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - B/op",
            "value": 1600,
            "unit": "B/op",
            "extra": "306914 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - allocs/op",
            "value": 42,
            "unit": "allocs/op",
            "extra": "306914 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric",
            "value": 5073,
            "unit": "ns/op\t    2208 B/op\t      58 allocs/op",
            "extra": "229868 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - ns/op",
            "value": 5073,
            "unit": "ns/op",
            "extra": "229868 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - B/op",
            "value": 2208,
            "unit": "B/op",
            "extra": "229868 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "229868 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode",
            "value": 6037,
            "unit": "ns/op\t    2128 B/op\t      56 allocs/op",
            "extra": "196668 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - ns/op",
            "value": 6037,
            "unit": "ns/op",
            "extra": "196668 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - B/op",
            "value": 2128,
            "unit": "B/op",
            "extra": "196668 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - allocs/op",
            "value": 56,
            "unit": "allocs/op",
            "extra": "196668 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode",
            "value": 7793,
            "unit": "ns/op\t    2769 B/op\t      73 allocs/op",
            "extra": "153375 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - ns/op",
            "value": 7793,
            "unit": "ns/op",
            "extra": "153375 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - B/op",
            "value": 2769,
            "unit": "B/op",
            "extra": "153375 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - allocs/op",
            "value": 73,
            "unit": "allocs/op",
            "extra": "153375 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4174,
            "unit": "ns/op\t    2054 B/op\t      58 allocs/op",
            "extra": "285199 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4174,
            "unit": "ns/op",
            "extra": "285199 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 2054,
            "unit": "B/op",
            "extra": "285199 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "285199 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1313,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "819877 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1313,
            "unit": "ns/op",
            "extra": "819877 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "819877 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "819877 times\n4 procs"
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
          "id": "aba1b1370491e4f1cdc0145825fbb1ebc8c18858",
          "message": "chore: Update dependency yaml to v2.6.1 (#53)\n\nThis PR contains the following updates:\n\n| Package | Change | Age | Adoption | Passing | Confidence |\n|---|---|---|---|---|---|\n| [yaml](https://eemeli.org/yaml/)\n([source](https://redirect.github.com/eemeli/yaml)) | [`2.6.0` ->\n`2.6.1`](https://renovatebot.com/diffs/npm/yaml/2.6.0/2.6.1) |\n[![age](https://developer.mend.io/api/mc/badges/age/npm/yaml/2.6.1?slim=true)](https://docs.renovatebot.com/merge-confidence/)\n|\n[![adoption](https://developer.mend.io/api/mc/badges/adoption/npm/yaml/2.6.1?slim=true)](https://docs.renovatebot.com/merge-confidence/)\n|\n[![passing](https://developer.mend.io/api/mc/badges/compatibility/npm/yaml/2.6.0/2.6.1?slim=true)](https://docs.renovatebot.com/merge-confidence/)\n|\n[![confidence](https://developer.mend.io/api/mc/badges/confidence/npm/yaml/2.6.0/2.6.1?slim=true)](https://docs.renovatebot.com/merge-confidence/)\n|\n\n---\n\n### Release Notes\n\n<details>\n<summary>eemeli/yaml (yaml)</summary>\n\n###\n[`v2.6.1`](https://redirect.github.com/eemeli/yaml/compare/v2.6.0...aa1898ae61605ea09bb79621d25ad5e7fd9b4217)\n\n[Compare\nSource](https://redirect.github.com/eemeli/yaml/compare/v2.6.0...v2.6.1)\n\n</details>\n\n---\n\n### Configuration\n\n📅 **Schedule**: Branch creation - \"after 10pm every weekday,before 5am\nevery weekday,every weekend\" (UTC), Automerge - At any time (no schedule\ndefined).\n\n🚦 **Automerge**: Enabled.\n\n♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the\nrebase/retry checkbox.\n\n🔕 **Ignore**: Close this PR and you won't be reminded about this update\nagain.\n\n---\n\n- [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check\nthis box\n\n---\n\nThis PR was generated by [Mend Renovate](https://mend.io/renovate/).\nView the [repository job\nlog](https://developer.mend.io/github/nobl9/govy).\n\n<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4xOS4wIiwidXBkYXRlZEluVmVyIjoiMzkuMTkuMCIsInRhcmdldEJyYW5jaCI6Im1haW4iLCJsYWJlbHMiOlsiZGVwZW5kZW5jaWVzIiwiamF2YXNjcmlwdCIsInJlbm92YXRlIl19-->\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2024-11-20T15:54:14+01:00",
          "tree_id": "ed739564b787bc14c8d0d3acb2f7fd6ecef1f2b7",
          "url": "https://github.com/nobl9/govy/commit/aba1b1370491e4f1cdc0145825fbb1ebc8c18858"
        },
        "date": 1732114592651,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 692.7,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1728606 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 692.7,
            "unit": "ns/op",
            "extra": "1728606 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1728606 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1728606 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 781.4,
            "unit": "ns/op\t     240 B/op\t       6 allocs/op",
            "extra": "1560763 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 781.4,
            "unit": "ns/op",
            "extra": "1560763 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "1560763 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1560763 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 863.4,
            "unit": "ns/op\t     400 B/op\t      10 allocs/op",
            "extra": "1446500 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 863.4,
            "unit": "ns/op",
            "extra": "1446500 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 400,
            "unit": "B/op",
            "extra": "1446500 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1446500 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 751.4,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1595619 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 751.4,
            "unit": "ns/op",
            "extra": "1595619 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1595619 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1595619 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 832.7,
            "unit": "ns/op\t     376 B/op\t      10 allocs/op",
            "extra": "1420760 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 832.7,
            "unit": "ns/op",
            "extra": "1420760 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 376,
            "unit": "B/op",
            "extra": "1420760 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1420760 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 749.4,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1611877 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 749.4,
            "unit": "ns/op",
            "extra": "1611877 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1611877 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1611877 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1229,
            "unit": "ns/op\t     520 B/op\t      18 allocs/op",
            "extra": "892441 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1229,
            "unit": "ns/op",
            "extra": "892441 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "892441 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "892441 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 178.2,
            "unit": "ns/op\t     144 B/op\t       4 allocs/op",
            "extra": "6865155 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 178.2,
            "unit": "ns/op",
            "extra": "6865155 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 144,
            "unit": "B/op",
            "extra": "6865155 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6865155 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1250,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "870075 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1250,
            "unit": "ns/op",
            "extra": "870075 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "870075 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "870075 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength",
            "value": 1026,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - ns/op",
            "value": 1026,
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
            "value": 1024,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1024,
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
            "value": 1255,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "866850 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1255,
            "unit": "ns/op",
            "extra": "866850 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "866850 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "866850 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 1030,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 1030,
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
            "value": 1051,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1051,
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
            "value": 1001,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1195326 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - ns/op",
            "value": 1001,
            "unit": "ns/op",
            "extra": "1195326 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1195326 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1195326 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf",
            "value": 1135,
            "unit": "ns/op\t     536 B/op\t      22 allocs/op",
            "extra": "943790 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1135,
            "unit": "ns/op",
            "extra": "943790 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 536,
            "unit": "B/op",
            "extra": "943790 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "943790 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 7247,
            "unit": "ns/op\t    3168 B/op\t      98 allocs/op",
            "extra": "161613 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 7247,
            "unit": "ns/op",
            "extra": "161613 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3168,
            "unit": "B/op",
            "extra": "161613 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "161613 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2571,
            "unit": "ns/op\t    1064 B/op\t      32 allocs/op",
            "extra": "448698 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2571,
            "unit": "ns/op",
            "extra": "448698 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1064,
            "unit": "B/op",
            "extra": "448698 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "448698 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1097,
            "unit": "ns/op\t     704 B/op\t      23 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1097,
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
            "value": 180.7,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "6590018 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 180.7,
            "unit": "ns/op",
            "extra": "6590018 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "6590018 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6590018 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 1507,
            "unit": "ns/op\t     644 B/op\t      20 allocs/op",
            "extra": "731372 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 1507,
            "unit": "ns/op",
            "extra": "731372 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 644,
            "unit": "B/op",
            "extra": "731372 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "731372 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 1567,
            "unit": "ns/op\t     692 B/op\t      20 allocs/op",
            "extra": "695767 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 1567,
            "unit": "ns/op",
            "extra": "695767 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 692,
            "unit": "B/op",
            "extra": "695767 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "695767 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 15414,
            "unit": "ns/op\t    5638 B/op\t     154 allocs/op",
            "extra": "77041 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 15414,
            "unit": "ns/op",
            "extra": "77041 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 5638,
            "unit": "B/op",
            "extra": "77041 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 154,
            "unit": "allocs/op",
            "extra": "77041 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4341,
            "unit": "ns/op\t    1552 B/op\t      41 allocs/op",
            "extra": "267248 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4341,
            "unit": "ns/op",
            "extra": "267248 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1552,
            "unit": "B/op",
            "extra": "267248 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "267248 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 16082,
            "unit": "ns/op\t   14085 B/op\t     217 allocs/op",
            "extra": "73338 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 16082,
            "unit": "ns/op",
            "extra": "73338 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 14085,
            "unit": "B/op",
            "extra": "73338 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "73338 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 8601,
            "unit": "ns/op\t    3408 B/op\t     138 allocs/op",
            "extra": "136326 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 8601,
            "unit": "ns/op",
            "extra": "136326 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3408,
            "unit": "B/op",
            "extra": "136326 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "136326 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 8965,
            "unit": "ns/op\t    5904 B/op\t      54 allocs/op",
            "extra": "132811 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 8965,
            "unit": "ns/op",
            "extra": "132811 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5904,
            "unit": "B/op",
            "extra": "132811 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "132811 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1078,
            "unit": "ns/op\t     752 B/op\t      25 allocs/op",
            "extra": "985821 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1078,
            "unit": "ns/op",
            "extra": "985821 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "985821 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "985821 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP",
            "value": 1575,
            "unit": "ns/op\t     824 B/op\t      31 allocs/op",
            "extra": "684080 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1575,
            "unit": "ns/op",
            "extra": "684080 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 824,
            "unit": "B/op",
            "extra": "684080 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "684080 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1689,
            "unit": "ns/op\t     896 B/op\t      32 allocs/op",
            "extra": "665613 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1689,
            "unit": "ns/op",
            "extra": "665613 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "665613 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "665613 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1858,
            "unit": "ns/op\t    1056 B/op\t      36 allocs/op",
            "extra": "608936 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1858,
            "unit": "ns/op",
            "extra": "608936 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "608936 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "608936 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 3021,
            "unit": "ns/op\t    1576 B/op\t      63 allocs/op",
            "extra": "377638 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 3021,
            "unit": "ns/op",
            "extra": "377638 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "377638 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "377638 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5608,
            "unit": "ns/op\t    3048 B/op\t     118 allocs/op",
            "extra": "209487 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5608,
            "unit": "ns/op",
            "extra": "209487 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 3048,
            "unit": "B/op",
            "extra": "209487 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "209487 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3609,
            "unit": "ns/op\t    2056 B/op\t      75 allocs/op",
            "extra": "318506 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3609,
            "unit": "ns/op",
            "extra": "318506 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 2056,
            "unit": "B/op",
            "extra": "318506 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "318506 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1172,
            "unit": "ns/op\t     616 B/op\t      23 allocs/op",
            "extra": "901568 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1172,
            "unit": "ns/op",
            "extra": "901568 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 616,
            "unit": "B/op",
            "extra": "901568 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "901568 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2390,
            "unit": "ns/op\t    1448 B/op\t      44 allocs/op",
            "extra": "472747 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2390,
            "unit": "ns/op",
            "extra": "472747 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "472747 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "472747 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2440,
            "unit": "ns/op\t    1576 B/op\t      46 allocs/op",
            "extra": "464174 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2440,
            "unit": "ns/op",
            "extra": "464174 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "464174 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "464174 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1356,
            "unit": "ns/op\t     784 B/op\t      24 allocs/op",
            "extra": "799161 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1356,
            "unit": "ns/op",
            "extra": "799161 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 784,
            "unit": "B/op",
            "extra": "799161 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "799161 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1338,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "791162 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1338,
            "unit": "ns/op",
            "extra": "791162 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "791162 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "791162 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1621,
            "unit": "ns/op\t     976 B/op\t      30 allocs/op",
            "extra": "686016 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1621,
            "unit": "ns/op",
            "extra": "686016 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "686016 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "686016 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 11750,
            "unit": "ns/op\t    3984 B/op\t      72 allocs/op",
            "extra": "100959 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 11750,
            "unit": "ns/op",
            "extra": "100959 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 3984,
            "unit": "B/op",
            "extra": "100959 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 72,
            "unit": "allocs/op",
            "extra": "100959 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 37245,
            "unit": "ns/op\t    7578 B/op\t      99 allocs/op",
            "extra": "32008 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 37245,
            "unit": "ns/op",
            "extra": "32008 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7578,
            "unit": "B/op",
            "extra": "32008 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "32008 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 38257,
            "unit": "ns/op\t    7833 B/op\t     108 allocs/op",
            "extra": "31680 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 38257,
            "unit": "ns/op",
            "extra": "31680 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7833,
            "unit": "B/op",
            "extra": "31680 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "31680 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 37615,
            "unit": "ns/op\t    7418 B/op\t     103 allocs/op",
            "extra": "31659 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 37615,
            "unit": "ns/op",
            "extra": "31659 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7418,
            "unit": "B/op",
            "extra": "31659 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "31659 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 16172,
            "unit": "ns/op\t    8721 B/op\t     217 allocs/op",
            "extra": "73410 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 16172,
            "unit": "ns/op",
            "extra": "73410 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8721,
            "unit": "B/op",
            "extra": "73410 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "73410 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 165926704,
            "unit": "ns/op\t334259405 B/op\t  281370 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 165926704,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334259405,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281370,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 46769,
            "unit": "ns/op\t   29417 B/op\t     614 allocs/op",
            "extra": "25274 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 46769,
            "unit": "ns/op",
            "extra": "25274 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 29417,
            "unit": "B/op",
            "extra": "25274 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "25274 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime",
            "value": 6300,
            "unit": "ns/op\t    3104 B/op\t      76 allocs/op",
            "extra": "187999 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - ns/op",
            "value": 6300,
            "unit": "ns/op",
            "extra": "187999 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - B/op",
            "value": 3104,
            "unit": "B/op",
            "extra": "187999 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - allocs/op",
            "value": 76,
            "unit": "allocs/op",
            "extra": "187999 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone",
            "value": 228541,
            "unit": "ns/op\t  337432 B/op\t     224 allocs/op",
            "extra": "4957 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - ns/op",
            "value": 228541,
            "unit": "ns/op",
            "extra": "4957 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - B/op",
            "value": 337432,
            "unit": "B/op",
            "extra": "4957 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "4957 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha",
            "value": 3715,
            "unit": "ns/op\t    1600 B/op\t      42 allocs/op",
            "extra": "307113 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - ns/op",
            "value": 3715,
            "unit": "ns/op",
            "extra": "307113 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - B/op",
            "value": 1600,
            "unit": "B/op",
            "extra": "307113 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - allocs/op",
            "value": 42,
            "unit": "allocs/op",
            "extra": "307113 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric",
            "value": 5092,
            "unit": "ns/op\t    2208 B/op\t      58 allocs/op",
            "extra": "230282 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - ns/op",
            "value": 5092,
            "unit": "ns/op",
            "extra": "230282 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - B/op",
            "value": 2208,
            "unit": "B/op",
            "extra": "230282 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "230282 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode",
            "value": 6069,
            "unit": "ns/op\t    2128 B/op\t      56 allocs/op",
            "extra": "193123 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - ns/op",
            "value": 6069,
            "unit": "ns/op",
            "extra": "193123 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - B/op",
            "value": 2128,
            "unit": "B/op",
            "extra": "193123 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - allocs/op",
            "value": 56,
            "unit": "allocs/op",
            "extra": "193123 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode",
            "value": 7792,
            "unit": "ns/op\t    2769 B/op\t      73 allocs/op",
            "extra": "151520 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - ns/op",
            "value": 7792,
            "unit": "ns/op",
            "extra": "151520 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - B/op",
            "value": 2769,
            "unit": "B/op",
            "extra": "151520 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - allocs/op",
            "value": 73,
            "unit": "allocs/op",
            "extra": "151520 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4131,
            "unit": "ns/op\t    2054 B/op\t      58 allocs/op",
            "extra": "278868 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4131,
            "unit": "ns/op",
            "extra": "278868 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 2054,
            "unit": "B/op",
            "extra": "278868 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "278868 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1342,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "838116 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1342,
            "unit": "ns/op",
            "extra": "838116 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "838116 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "838116 times\n4 procs"
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
          "id": "d65c01e6711acf809ba7a71c1d676275dd5449d6",
          "message": "feat: Add EqualProperties rule (#54)\n\n## Release Notes\r\n\r\nAdded `rules.EqualProperties` rule which helps ensure selected\r\nproperties are equal.\r\nThe equality check is performed via a configurable function.\r\nTwo builtin functions are provided out of the box: `rules.CompareFunc`\r\nwhich operates on `comparable` types and `rules.CompareDeepEqualFunc`\r\nwhich uses `reflect.DeepEqual` and operates on any type.",
          "timestamp": "2024-11-20T18:07:10+01:00",
          "tree_id": "d4cd36fe4638f597c34d672a0b78794dea89e13a",
          "url": "https://github.com/nobl9/govy/commit/d65c01e6711acf809ba7a71c1d676275dd5449d6"
        },
        "date": 1732122597951,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 688.2,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1743811 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 688.2,
            "unit": "ns/op",
            "extra": "1743811 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1743811 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1743811 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 772.4,
            "unit": "ns/op\t     240 B/op\t       6 allocs/op",
            "extra": "1572098 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 772.4,
            "unit": "ns/op",
            "extra": "1572098 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "1572098 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1572098 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 817,
            "unit": "ns/op\t     400 B/op\t      10 allocs/op",
            "extra": "1463913 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 817,
            "unit": "ns/op",
            "extra": "1463913 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 400,
            "unit": "B/op",
            "extra": "1463913 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1463913 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 742.6,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1638024 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 742.6,
            "unit": "ns/op",
            "extra": "1638024 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1638024 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1638024 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 823.9,
            "unit": "ns/op\t     376 B/op\t      10 allocs/op",
            "extra": "1471941 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 823.9,
            "unit": "ns/op",
            "extra": "1471941 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 376,
            "unit": "B/op",
            "extra": "1471941 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1471941 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 785.2,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1575433 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 785.2,
            "unit": "ns/op",
            "extra": "1575433 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1575433 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1575433 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties",
            "value": 6800,
            "unit": "ns/op\t    2720 B/op\t      83 allocs/op",
            "extra": "172718 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - ns/op",
            "value": 6800,
            "unit": "ns/op",
            "extra": "172718 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - B/op",
            "value": 2720,
            "unit": "B/op",
            "extra": "172718 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - allocs/op",
            "value": 83,
            "unit": "allocs/op",
            "extra": "172718 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1266,
            "unit": "ns/op\t     520 B/op\t      18 allocs/op",
            "extra": "858853 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1266,
            "unit": "ns/op",
            "extra": "858853 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "858853 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "858853 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 173.4,
            "unit": "ns/op\t     144 B/op\t       4 allocs/op",
            "extra": "6867424 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 173.4,
            "unit": "ns/op",
            "extra": "6867424 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 144,
            "unit": "B/op",
            "extra": "6867424 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6867424 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1262,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "866468 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1262,
            "unit": "ns/op",
            "extra": "866468 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "866468 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "866468 times\n4 procs"
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
            "value": 1064,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "998946 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1064,
            "unit": "ns/op",
            "extra": "998946 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "998946 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "998946 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength",
            "value": 1265,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "877209 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1265,
            "unit": "ns/op",
            "extra": "877209 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "877209 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "877209 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 1031,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 1031,
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
            "value": 1024,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - ns/op",
            "value": 1024,
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
            "value": 1125,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "993260 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1125,
            "unit": "ns/op",
            "extra": "993260 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - B/op",
            "value": 560,
            "unit": "B/op",
            "extra": "993260 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "993260 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength",
            "value": 1016,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - ns/op",
            "value": 1016,
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
            "value": 1007,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - ns/op",
            "value": 1007,
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
            "value": 1120,
            "unit": "ns/op\t     536 B/op\t      22 allocs/op",
            "extra": "952574 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1120,
            "unit": "ns/op",
            "extra": "952574 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 536,
            "unit": "B/op",
            "extra": "952574 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "952574 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 7525,
            "unit": "ns/op\t    3168 B/op\t      98 allocs/op",
            "extra": "156386 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 7525,
            "unit": "ns/op",
            "extra": "156386 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3168,
            "unit": "B/op",
            "extra": "156386 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "156386 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2868,
            "unit": "ns/op\t    1064 B/op\t      32 allocs/op",
            "extra": "418083 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2868,
            "unit": "ns/op",
            "extra": "418083 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1064,
            "unit": "B/op",
            "extra": "418083 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "418083 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1089,
            "unit": "ns/op\t     704 B/op\t      23 allocs/op",
            "extra": "993636 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1089,
            "unit": "ns/op",
            "extra": "993636 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - B/op",
            "value": 704,
            "unit": "B/op",
            "extra": "993636 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "993636 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty",
            "value": 184.1,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "6624748 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 184.1,
            "unit": "ns/op",
            "extra": "6624748 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "6624748 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6624748 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 1521,
            "unit": "ns/op\t     644 B/op\t      20 allocs/op",
            "extra": "788988 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 1521,
            "unit": "ns/op",
            "extra": "788988 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 644,
            "unit": "B/op",
            "extra": "788988 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "788988 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 1481,
            "unit": "ns/op\t     693 B/op\t      20 allocs/op",
            "extra": "810010 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 1481,
            "unit": "ns/op",
            "extra": "810010 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 693,
            "unit": "B/op",
            "extra": "810010 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "810010 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 15447,
            "unit": "ns/op\t    5643 B/op\t     154 allocs/op",
            "extra": "77380 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 15447,
            "unit": "ns/op",
            "extra": "77380 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 5643,
            "unit": "B/op",
            "extra": "77380 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 154,
            "unit": "allocs/op",
            "extra": "77380 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4410,
            "unit": "ns/op\t    1552 B/op\t      41 allocs/op",
            "extra": "265221 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4410,
            "unit": "ns/op",
            "extra": "265221 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1552,
            "unit": "B/op",
            "extra": "265221 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "265221 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 16233,
            "unit": "ns/op\t   14085 B/op\t     217 allocs/op",
            "extra": "72972 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 16233,
            "unit": "ns/op",
            "extra": "72972 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 14085,
            "unit": "B/op",
            "extra": "72972 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "72972 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 8804,
            "unit": "ns/op\t    3408 B/op\t     138 allocs/op",
            "extra": "132140 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 8804,
            "unit": "ns/op",
            "extra": "132140 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3408,
            "unit": "B/op",
            "extra": "132140 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "132140 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 9152,
            "unit": "ns/op\t    5904 B/op\t      54 allocs/op",
            "extra": "129444 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 9152,
            "unit": "ns/op",
            "extra": "129444 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5904,
            "unit": "B/op",
            "extra": "129444 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "129444 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1074,
            "unit": "ns/op\t     752 B/op\t      25 allocs/op",
            "extra": "980980 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1074,
            "unit": "ns/op",
            "extra": "980980 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "980980 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "980980 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP",
            "value": 1650,
            "unit": "ns/op\t     824 B/op\t      31 allocs/op",
            "extra": "685480 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1650,
            "unit": "ns/op",
            "extra": "685480 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 824,
            "unit": "B/op",
            "extra": "685480 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "685480 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1708,
            "unit": "ns/op\t     896 B/op\t      32 allocs/op",
            "extra": "666523 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1708,
            "unit": "ns/op",
            "extra": "666523 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "666523 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "666523 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1846,
            "unit": "ns/op\t    1056 B/op\t      36 allocs/op",
            "extra": "615961 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1846,
            "unit": "ns/op",
            "extra": "615961 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "615961 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "615961 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 3032,
            "unit": "ns/op\t    1576 B/op\t      63 allocs/op",
            "extra": "379348 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 3032,
            "unit": "ns/op",
            "extra": "379348 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "379348 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "379348 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5565,
            "unit": "ns/op\t    3048 B/op\t     118 allocs/op",
            "extra": "210589 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5565,
            "unit": "ns/op",
            "extra": "210589 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 3048,
            "unit": "B/op",
            "extra": "210589 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "210589 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3563,
            "unit": "ns/op\t    2056 B/op\t      75 allocs/op",
            "extra": "324330 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3563,
            "unit": "ns/op",
            "extra": "324330 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 2056,
            "unit": "B/op",
            "extra": "324330 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "324330 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1174,
            "unit": "ns/op\t     616 B/op\t      23 allocs/op",
            "extra": "923468 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1174,
            "unit": "ns/op",
            "extra": "923468 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 616,
            "unit": "B/op",
            "extra": "923468 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "923468 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2336,
            "unit": "ns/op\t    1448 B/op\t      44 allocs/op",
            "extra": "475017 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2336,
            "unit": "ns/op",
            "extra": "475017 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "475017 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "475017 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2414,
            "unit": "ns/op\t    1576 B/op\t      46 allocs/op",
            "extra": "469494 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2414,
            "unit": "ns/op",
            "extra": "469494 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "469494 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "469494 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1345,
            "unit": "ns/op\t     784 B/op\t      24 allocs/op",
            "extra": "830575 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1345,
            "unit": "ns/op",
            "extra": "830575 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 784,
            "unit": "B/op",
            "extra": "830575 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "830575 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1347,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "828212 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1347,
            "unit": "ns/op",
            "extra": "828212 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "828212 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "828212 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1623,
            "unit": "ns/op\t     976 B/op\t      30 allocs/op",
            "extra": "699223 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1623,
            "unit": "ns/op",
            "extra": "699223 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "699223 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "699223 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 12389,
            "unit": "ns/op\t    3984 B/op\t      72 allocs/op",
            "extra": "97940 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 12389,
            "unit": "ns/op",
            "extra": "97940 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 3984,
            "unit": "B/op",
            "extra": "97940 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 72,
            "unit": "allocs/op",
            "extra": "97940 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 37371,
            "unit": "ns/op\t    7577 B/op\t      99 allocs/op",
            "extra": "32122 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 37371,
            "unit": "ns/op",
            "extra": "32122 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7577,
            "unit": "B/op",
            "extra": "32122 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "32122 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 37713,
            "unit": "ns/op\t    7834 B/op\t     108 allocs/op",
            "extra": "31845 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 37713,
            "unit": "ns/op",
            "extra": "31845 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7834,
            "unit": "B/op",
            "extra": "31845 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "31845 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 37372,
            "unit": "ns/op\t    7449 B/op\t     103 allocs/op",
            "extra": "31899 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 37372,
            "unit": "ns/op",
            "extra": "31899 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7449,
            "unit": "B/op",
            "extra": "31899 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "31899 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 16810,
            "unit": "ns/op\t    8721 B/op\t     217 allocs/op",
            "extra": "74935 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 16810,
            "unit": "ns/op",
            "extra": "74935 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8721,
            "unit": "B/op",
            "extra": "74935 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "74935 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 165019116,
            "unit": "ns/op\t334259001 B/op\t  281367 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 165019116,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334259001,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281367,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 47309,
            "unit": "ns/op\t   29417 B/op\t     614 allocs/op",
            "extra": "25660 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 47309,
            "unit": "ns/op",
            "extra": "25660 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 29417,
            "unit": "B/op",
            "extra": "25660 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "25660 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime",
            "value": 6419,
            "unit": "ns/op\t    3104 B/op\t      76 allocs/op",
            "extra": "182862 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - ns/op",
            "value": 6419,
            "unit": "ns/op",
            "extra": "182862 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - B/op",
            "value": 3104,
            "unit": "B/op",
            "extra": "182862 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - allocs/op",
            "value": 76,
            "unit": "allocs/op",
            "extra": "182862 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone",
            "value": 233048,
            "unit": "ns/op\t  337429 B/op\t     224 allocs/op",
            "extra": "4898 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - ns/op",
            "value": 233048,
            "unit": "ns/op",
            "extra": "4898 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - B/op",
            "value": 337429,
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
            "name": "BenchmarkStringAlpha",
            "value": 3802,
            "unit": "ns/op\t    1600 B/op\t      42 allocs/op",
            "extra": "309054 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - ns/op",
            "value": 3802,
            "unit": "ns/op",
            "extra": "309054 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - B/op",
            "value": 1600,
            "unit": "B/op",
            "extra": "309054 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - allocs/op",
            "value": 42,
            "unit": "allocs/op",
            "extra": "309054 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric",
            "value": 5278,
            "unit": "ns/op\t    2208 B/op\t      58 allocs/op",
            "extra": "226240 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - ns/op",
            "value": 5278,
            "unit": "ns/op",
            "extra": "226240 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - B/op",
            "value": 2208,
            "unit": "B/op",
            "extra": "226240 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "226240 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode",
            "value": 6128,
            "unit": "ns/op\t    2128 B/op\t      56 allocs/op",
            "extra": "192421 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - ns/op",
            "value": 6128,
            "unit": "ns/op",
            "extra": "192421 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - B/op",
            "value": 2128,
            "unit": "B/op",
            "extra": "192421 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - allocs/op",
            "value": 56,
            "unit": "allocs/op",
            "extra": "192421 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode",
            "value": 7897,
            "unit": "ns/op\t    2769 B/op\t      73 allocs/op",
            "extra": "147183 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - ns/op",
            "value": 7897,
            "unit": "ns/op",
            "extra": "147183 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - B/op",
            "value": 2769,
            "unit": "B/op",
            "extra": "147183 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - allocs/op",
            "value": 73,
            "unit": "allocs/op",
            "extra": "147183 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4191,
            "unit": "ns/op\t    2054 B/op\t      58 allocs/op",
            "extra": "276774 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4191,
            "unit": "ns/op",
            "extra": "276774 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 2054,
            "unit": "B/op",
            "extra": "276774 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "276774 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1315,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "844737 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1315,
            "unit": "ns/op",
            "extra": "844737 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "844737 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "844737 times\n4 procs"
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
          "id": "c1552c3a2e836abebfb7f63fe952949869e675c2",
          "message": "chore: Update dependency markdownlint-cli to v0.43.0 (#55)\n\nThis PR contains the following updates:\n\n| Package | Change | Age | Adoption | Passing | Confidence |\n|---|---|---|---|---|---|\n|\n[markdownlint-cli](https://redirect.github.com/igorshubovych/markdownlint-cli)\n| [`0.42.0` ->\n`0.43.0`](https://renovatebot.com/diffs/npm/markdownlint-cli/0.42.0/0.43.0)\n|\n[![age](https://developer.mend.io/api/mc/badges/age/npm/markdownlint-cli/0.43.0?slim=true)](https://docs.renovatebot.com/merge-confidence/)\n|\n[![adoption](https://developer.mend.io/api/mc/badges/adoption/npm/markdownlint-cli/0.43.0?slim=true)](https://docs.renovatebot.com/merge-confidence/)\n|\n[![passing](https://developer.mend.io/api/mc/badges/compatibility/npm/markdownlint-cli/0.42.0/0.43.0?slim=true)](https://docs.renovatebot.com/merge-confidence/)\n|\n[![confidence](https://developer.mend.io/api/mc/badges/confidence/npm/markdownlint-cli/0.42.0/0.43.0?slim=true)](https://docs.renovatebot.com/merge-confidence/)\n|\n\n---\n\n### Release Notes\n\n<details>\n<summary>igorshubovych/markdownlint-cli (markdownlint-cli)</summary>\n\n###\n[`v0.43.0`](https://redirect.github.com/igorshubovych/markdownlint-cli/compare/v0.42.0...v0.43.0)\n\n[Compare\nSource](https://redirect.github.com/igorshubovych/markdownlint-cli/compare/v0.42.0...v0.43.0)\n\n</details>\n\n---\n\n### Configuration\n\n📅 **Schedule**: Branch creation - \"after 10pm every weekday,before 5am\nevery weekday,every weekend\" (UTC), Automerge - At any time (no schedule\ndefined).\n\n🚦 **Automerge**: Enabled.\n\n♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the\nrebase/retry checkbox.\n\n🔕 **Ignore**: Close this PR and you won't be reminded about this update\nagain.\n\n---\n\n- [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check\nthis box\n\n---\n\nThis PR was generated by [Mend Renovate](https://mend.io/renovate/).\nView the [repository job\nlog](https://developer.mend.io/github/nobl9/govy).\n\n<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4xOS4wIiwidXBkYXRlZEluVmVyIjoiMzkuMTkuMCIsInRhcmdldEJyYW5jaCI6Im1haW4iLCJsYWJlbHMiOlsiZGVwZW5kZW5jaWVzIiwiamF2YXNjcmlwdCIsInJlbm92YXRlIl19-->\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2024-11-24T16:15:01+01:00",
          "tree_id": "2c96652301dec8d827e94ffd523cd85b30687a35",
          "url": "https://github.com/nobl9/govy/commit/c1552c3a2e836abebfb7f63fe952949869e675c2"
        },
        "date": 1732461450311,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 677.3,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1770109 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 677.3,
            "unit": "ns/op",
            "extra": "1770109 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1770109 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1770109 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 759.2,
            "unit": "ns/op\t     240 B/op\t       6 allocs/op",
            "extra": "1581816 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 759.2,
            "unit": "ns/op",
            "extra": "1581816 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "1581816 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1581816 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 825.5,
            "unit": "ns/op\t     400 B/op\t      10 allocs/op",
            "extra": "1392824 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 825.5,
            "unit": "ns/op",
            "extra": "1392824 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 400,
            "unit": "B/op",
            "extra": "1392824 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1392824 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 726.4,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1658736 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 726.4,
            "unit": "ns/op",
            "extra": "1658736 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1658736 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1658736 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 799.5,
            "unit": "ns/op\t     376 B/op\t      10 allocs/op",
            "extra": "1501706 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 799.5,
            "unit": "ns/op",
            "extra": "1501706 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 376,
            "unit": "B/op",
            "extra": "1501706 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1501706 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 742.8,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1599380 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 742.8,
            "unit": "ns/op",
            "extra": "1599380 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1599380 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1599380 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties",
            "value": 6675,
            "unit": "ns/op\t    2720 B/op\t      83 allocs/op",
            "extra": "176049 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - ns/op",
            "value": 6675,
            "unit": "ns/op",
            "extra": "176049 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - B/op",
            "value": 2720,
            "unit": "B/op",
            "extra": "176049 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - allocs/op",
            "value": 83,
            "unit": "allocs/op",
            "extra": "176049 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1253,
            "unit": "ns/op\t     520 B/op\t      18 allocs/op",
            "extra": "840097 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1253,
            "unit": "ns/op",
            "extra": "840097 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "840097 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "840097 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 171.5,
            "unit": "ns/op\t     144 B/op\t       4 allocs/op",
            "extra": "6975261 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 171.5,
            "unit": "ns/op",
            "extra": "6975261 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 144,
            "unit": "B/op",
            "extra": "6975261 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6975261 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1247,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "873000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1247,
            "unit": "ns/op",
            "extra": "873000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "873000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "873000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength",
            "value": 1012,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - ns/op",
            "value": 1012,
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
            "value": 1035,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1035,
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
            "value": 1246,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "866796 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1246,
            "unit": "ns/op",
            "extra": "866796 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "866796 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "866796 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 1032,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 1032,
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
            "value": 1001,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - ns/op",
            "value": 1001,
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
            "value": 1055,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1055,
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
            "value": 994.9,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1206369 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - ns/op",
            "value": 994.9,
            "unit": "ns/op",
            "extra": "1206369 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1206369 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1206369 times\n4 procs"
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
            "value": 1135,
            "unit": "ns/op\t     536 B/op\t      22 allocs/op",
            "extra": "967490 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1135,
            "unit": "ns/op",
            "extra": "967490 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 536,
            "unit": "B/op",
            "extra": "967490 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "967490 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 7441,
            "unit": "ns/op\t    3168 B/op\t      98 allocs/op",
            "extra": "157993 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 7441,
            "unit": "ns/op",
            "extra": "157993 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3168,
            "unit": "B/op",
            "extra": "157993 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "157993 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2697,
            "unit": "ns/op\t    1064 B/op\t      32 allocs/op",
            "extra": "420968 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2697,
            "unit": "ns/op",
            "extra": "420968 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1064,
            "unit": "B/op",
            "extra": "420968 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "420968 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1078,
            "unit": "ns/op\t     704 B/op\t      23 allocs/op",
            "extra": "986816 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1078,
            "unit": "ns/op",
            "extra": "986816 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - B/op",
            "value": 704,
            "unit": "B/op",
            "extra": "986816 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "986816 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty",
            "value": 180.6,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "6587101 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 180.6,
            "unit": "ns/op",
            "extra": "6587101 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "6587101 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6587101 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 1483,
            "unit": "ns/op\t     644 B/op\t      20 allocs/op",
            "extra": "747198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 1483,
            "unit": "ns/op",
            "extra": "747198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 644,
            "unit": "B/op",
            "extra": "747198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "747198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 1474,
            "unit": "ns/op\t     693 B/op\t      20 allocs/op",
            "extra": "736090 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 1474,
            "unit": "ns/op",
            "extra": "736090 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 693,
            "unit": "B/op",
            "extra": "736090 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "736090 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 15261,
            "unit": "ns/op\t    5638 B/op\t     154 allocs/op",
            "extra": "77325 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 15261,
            "unit": "ns/op",
            "extra": "77325 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 5638,
            "unit": "B/op",
            "extra": "77325 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 154,
            "unit": "allocs/op",
            "extra": "77325 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4386,
            "unit": "ns/op\t    1552 B/op\t      41 allocs/op",
            "extra": "266972 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4386,
            "unit": "ns/op",
            "extra": "266972 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1552,
            "unit": "B/op",
            "extra": "266972 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "266972 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 16001,
            "unit": "ns/op\t   14085 B/op\t     217 allocs/op",
            "extra": "74317 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 16001,
            "unit": "ns/op",
            "extra": "74317 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 14085,
            "unit": "B/op",
            "extra": "74317 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "74317 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 8656,
            "unit": "ns/op\t    3408 B/op\t     138 allocs/op",
            "extra": "136800 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 8656,
            "unit": "ns/op",
            "extra": "136800 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3408,
            "unit": "B/op",
            "extra": "136800 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "136800 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 9117,
            "unit": "ns/op\t    5904 B/op\t      54 allocs/op",
            "extra": "127962 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 9117,
            "unit": "ns/op",
            "extra": "127962 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5904,
            "unit": "B/op",
            "extra": "127962 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "127962 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1061,
            "unit": "ns/op\t     752 B/op\t      25 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1061,
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
            "value": 1579,
            "unit": "ns/op\t     824 B/op\t      31 allocs/op",
            "extra": "703275 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1579,
            "unit": "ns/op",
            "extra": "703275 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 824,
            "unit": "B/op",
            "extra": "703275 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "703275 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1674,
            "unit": "ns/op\t     896 B/op\t      32 allocs/op",
            "extra": "673622 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1674,
            "unit": "ns/op",
            "extra": "673622 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "673622 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "673622 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1830,
            "unit": "ns/op\t    1056 B/op\t      36 allocs/op",
            "extra": "611989 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1830,
            "unit": "ns/op",
            "extra": "611989 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "611989 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "611989 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 2987,
            "unit": "ns/op\t    1576 B/op\t      63 allocs/op",
            "extra": "381896 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 2987,
            "unit": "ns/op",
            "extra": "381896 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "381896 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "381896 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5508,
            "unit": "ns/op\t    3048 B/op\t     118 allocs/op",
            "extra": "213087 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5508,
            "unit": "ns/op",
            "extra": "213087 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 3048,
            "unit": "B/op",
            "extra": "213087 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "213087 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3520,
            "unit": "ns/op\t    2056 B/op\t      75 allocs/op",
            "extra": "328544 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3520,
            "unit": "ns/op",
            "extra": "328544 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 2056,
            "unit": "B/op",
            "extra": "328544 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "328544 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1131,
            "unit": "ns/op\t     616 B/op\t      23 allocs/op",
            "extra": "958032 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1131,
            "unit": "ns/op",
            "extra": "958032 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 616,
            "unit": "B/op",
            "extra": "958032 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "958032 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2304,
            "unit": "ns/op\t    1448 B/op\t      44 allocs/op",
            "extra": "502672 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2304,
            "unit": "ns/op",
            "extra": "502672 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "502672 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "502672 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2384,
            "unit": "ns/op\t    1576 B/op\t      46 allocs/op",
            "extra": "479528 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2384,
            "unit": "ns/op",
            "extra": "479528 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "479528 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "479528 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1332,
            "unit": "ns/op\t     784 B/op\t      24 allocs/op",
            "extra": "794050 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1332,
            "unit": "ns/op",
            "extra": "794050 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 784,
            "unit": "B/op",
            "extra": "794050 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "794050 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1320,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "826156 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1320,
            "unit": "ns/op",
            "extra": "826156 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "826156 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "826156 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1595,
            "unit": "ns/op\t     976 B/op\t      30 allocs/op",
            "extra": "687664 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1595,
            "unit": "ns/op",
            "extra": "687664 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "687664 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "687664 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 12031,
            "unit": "ns/op\t    3984 B/op\t      72 allocs/op",
            "extra": "99142 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 12031,
            "unit": "ns/op",
            "extra": "99142 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 3984,
            "unit": "B/op",
            "extra": "99142 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 72,
            "unit": "allocs/op",
            "extra": "99142 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 36960,
            "unit": "ns/op\t    7529 B/op\t      99 allocs/op",
            "extra": "32251 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 36960,
            "unit": "ns/op",
            "extra": "32251 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7529,
            "unit": "B/op",
            "extra": "32251 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "32251 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 37500,
            "unit": "ns/op\t    7834 B/op\t     108 allocs/op",
            "extra": "31992 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 37500,
            "unit": "ns/op",
            "extra": "31992 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7834,
            "unit": "B/op",
            "extra": "31992 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "31992 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 36698,
            "unit": "ns/op\t    7417 B/op\t     103 allocs/op",
            "extra": "32355 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 36698,
            "unit": "ns/op",
            "extra": "32355 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7417,
            "unit": "B/op",
            "extra": "32355 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "32355 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 15762,
            "unit": "ns/op\t    8721 B/op\t     217 allocs/op",
            "extra": "75702 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 15762,
            "unit": "ns/op",
            "extra": "75702 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8721,
            "unit": "B/op",
            "extra": "75702 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "75702 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 164488375,
            "unit": "ns/op\t334259686 B/op\t  281372 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 164488375,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334259686,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281372,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 46402,
            "unit": "ns/op\t   29417 B/op\t     614 allocs/op",
            "extra": "25664 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 46402,
            "unit": "ns/op",
            "extra": "25664 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 29417,
            "unit": "B/op",
            "extra": "25664 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "25664 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime",
            "value": 6223,
            "unit": "ns/op\t    3104 B/op\t      76 allocs/op",
            "extra": "187785 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - ns/op",
            "value": 6223,
            "unit": "ns/op",
            "extra": "187785 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - B/op",
            "value": 3104,
            "unit": "B/op",
            "extra": "187785 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - allocs/op",
            "value": 76,
            "unit": "allocs/op",
            "extra": "187785 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone",
            "value": 224810,
            "unit": "ns/op\t  337434 B/op\t     224 allocs/op",
            "extra": "5120 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - ns/op",
            "value": 224810,
            "unit": "ns/op",
            "extra": "5120 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - B/op",
            "value": 337434,
            "unit": "B/op",
            "extra": "5120 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "5120 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha",
            "value": 3739,
            "unit": "ns/op\t    1600 B/op\t      42 allocs/op",
            "extra": "307646 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - ns/op",
            "value": 3739,
            "unit": "ns/op",
            "extra": "307646 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - B/op",
            "value": 1600,
            "unit": "B/op",
            "extra": "307646 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - allocs/op",
            "value": 42,
            "unit": "allocs/op",
            "extra": "307646 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric",
            "value": 5117,
            "unit": "ns/op\t    2208 B/op\t      58 allocs/op",
            "extra": "229172 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - ns/op",
            "value": 5117,
            "unit": "ns/op",
            "extra": "229172 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - B/op",
            "value": 2208,
            "unit": "B/op",
            "extra": "229172 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "229172 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode",
            "value": 6044,
            "unit": "ns/op\t    2128 B/op\t      56 allocs/op",
            "extra": "194886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - ns/op",
            "value": 6044,
            "unit": "ns/op",
            "extra": "194886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - B/op",
            "value": 2128,
            "unit": "B/op",
            "extra": "194886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - allocs/op",
            "value": 56,
            "unit": "allocs/op",
            "extra": "194886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode",
            "value": 7846,
            "unit": "ns/op\t    2769 B/op\t      73 allocs/op",
            "extra": "150646 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - ns/op",
            "value": 7846,
            "unit": "ns/op",
            "extra": "150646 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - B/op",
            "value": 2769,
            "unit": "B/op",
            "extra": "150646 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - allocs/op",
            "value": 73,
            "unit": "allocs/op",
            "extra": "150646 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4150,
            "unit": "ns/op\t    2054 B/op\t      58 allocs/op",
            "extra": "283426 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4150,
            "unit": "ns/op",
            "extra": "283426 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 2054,
            "unit": "B/op",
            "extra": "283426 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "283426 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1309,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "845696 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1309,
            "unit": "ns/op",
            "extra": "845696 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "845696 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "845696 times\n4 procs"
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
          "id": "10fa133a5b704e080d78ace0412ffa9ac57b5018",
          "message": "chore: Remove details from govy.RuleError (#60)\n\n## Motivation\r\n\r\nWe're moving towards a more unified error message creation. The goal is\r\nto allow users to manipulate the error message however they see fit.\r\nWith that in mind, `govy.RuleError` should have its `Message` field\r\npopulated with a created error that is ready to be used as is and passed\r\ndownstream.\r\n\r\n## Related Changes\r\n\r\n- https://github.com/nobl9/govy/pull/10\r\n- https://github.com/nobl9/govy/pull/59\r\n\r\n## Breaking Changes\r\n\r\nRemoved `govy.RuleError.Details` field, the error's details are now part\r\nof the `govy.RuleError.Message` and not just the\r\n`govy.RuleError.Error()` output.",
          "timestamp": "2024-12-20T22:59:53+01:00",
          "tree_id": "5ce88fdda5076840fdd0af684a1bec6b1ddfff97",
          "url": "https://github.com/nobl9/govy/commit/10fa133a5b704e080d78ace0412ffa9ac57b5018"
        },
        "date": 1736767031323,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 719.9,
            "unit": "ns/op\t     208 B/op\t       6 allocs/op",
            "extra": "1711020 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 719.9,
            "unit": "ns/op",
            "extra": "1711020 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 208,
            "unit": "B/op",
            "extra": "1711020 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1711020 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 812.7,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1471945 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 812.7,
            "unit": "ns/op",
            "extra": "1471945 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1471945 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1471945 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 850.7,
            "unit": "ns/op\t     368 B/op\t      10 allocs/op",
            "extra": "1396033 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 850.7,
            "unit": "ns/op",
            "extra": "1396033 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1396033 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1396033 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 759.1,
            "unit": "ns/op\t     352 B/op\t       8 allocs/op",
            "extra": "1582238 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 759.1,
            "unit": "ns/op",
            "extra": "1582238 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 352,
            "unit": "B/op",
            "extra": "1582238 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1582238 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 838,
            "unit": "ns/op\t     344 B/op\t      10 allocs/op",
            "extra": "1441666 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 838,
            "unit": "ns/op",
            "extra": "1441666 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 344,
            "unit": "B/op",
            "extra": "1441666 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1441666 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 766.9,
            "unit": "ns/op\t     352 B/op\t       8 allocs/op",
            "extra": "1560076 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 766.9,
            "unit": "ns/op",
            "extra": "1560076 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 352,
            "unit": "B/op",
            "extra": "1560076 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1560076 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties",
            "value": 6717,
            "unit": "ns/op\t    2672 B/op\t      83 allocs/op",
            "extra": "172063 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - ns/op",
            "value": 6717,
            "unit": "ns/op",
            "extra": "172063 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - B/op",
            "value": 2672,
            "unit": "B/op",
            "extra": "172063 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - allocs/op",
            "value": 83,
            "unit": "allocs/op",
            "extra": "172063 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1235,
            "unit": "ns/op\t     488 B/op\t      18 allocs/op",
            "extra": "900105 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1235,
            "unit": "ns/op",
            "extra": "900105 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 488,
            "unit": "B/op",
            "extra": "900105 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "900105 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 170.7,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "7010448 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 170.7,
            "unit": "ns/op",
            "extra": "7010448 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "7010448 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "7010448 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1332,
            "unit": "ns/op\t     608 B/op\t      16 allocs/op",
            "extra": "818167 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1332,
            "unit": "ns/op",
            "extra": "818167 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 608,
            "unit": "B/op",
            "extra": "818167 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "818167 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength",
            "value": 1046,
            "unit": "ns/op\t     448 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - ns/op",
            "value": 1046,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - B/op",
            "value": 448,
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
            "value": 1079,
            "unit": "ns/op\t     448 B/op\t      12 allocs/op",
            "extra": "983202 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1079,
            "unit": "ns/op",
            "extra": "983202 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "983202 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "983202 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength",
            "value": 1309,
            "unit": "ns/op\t     608 B/op\t      16 allocs/op",
            "extra": "830161 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1309,
            "unit": "ns/op",
            "extra": "830161 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 608,
            "unit": "B/op",
            "extra": "830161 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "830161 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 1034,
            "unit": "ns/op\t     512 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 1034,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - B/op",
            "value": 512,
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
            "value": 1060,
            "unit": "ns/op\t     512 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - ns/op",
            "value": 1060,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - B/op",
            "value": 512,
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
            "value": 1116,
            "unit": "ns/op\t     528 B/op\t      14 allocs/op",
            "extra": "952838 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1116,
            "unit": "ns/op",
            "extra": "952838 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - B/op",
            "value": 528,
            "unit": "B/op",
            "extra": "952838 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "952838 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength",
            "value": 1039,
            "unit": "ns/op\t     512 B/op\t      12 allocs/op",
            "extra": "997462 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - ns/op",
            "value": 1039,
            "unit": "ns/op",
            "extra": "997462 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - B/op",
            "value": 512,
            "unit": "B/op",
            "extra": "997462 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "997462 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength",
            "value": 1049,
            "unit": "ns/op\t     512 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - ns/op",
            "value": 1049,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - B/op",
            "value": 512,
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
            "value": 1062,
            "unit": "ns/op\t     520 B/op\t      22 allocs/op",
            "extra": "986654 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1062,
            "unit": "ns/op",
            "extra": "986654 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "986654 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "986654 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 7308,
            "unit": "ns/op\t    3088 B/op\t      98 allocs/op",
            "extra": "161282 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 7308,
            "unit": "ns/op",
            "extra": "161282 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3088,
            "unit": "B/op",
            "extra": "161282 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "161282 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2643,
            "unit": "ns/op\t    1048 B/op\t      32 allocs/op",
            "extra": "432842 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2643,
            "unit": "ns/op",
            "extra": "432842 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "432842 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "432842 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1051,
            "unit": "ns/op\t     608 B/op\t      23 allocs/op",
            "extra": "1143453 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1051,
            "unit": "ns/op",
            "extra": "1143453 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - B/op",
            "value": 608,
            "unit": "B/op",
            "extra": "1143453 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "1143453 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty",
            "value": 176.8,
            "unit": "ns/op\t     112 B/op\t       4 allocs/op",
            "extra": "6735470 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 176.8,
            "unit": "ns/op",
            "extra": "6735470 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "6735470 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6735470 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 1475,
            "unit": "ns/op\t     611 B/op\t      20 allocs/op",
            "extra": "752406 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 1475,
            "unit": "ns/op",
            "extra": "752406 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 611,
            "unit": "B/op",
            "extra": "752406 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "752406 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 1490,
            "unit": "ns/op\t     660 B/op\t      20 allocs/op",
            "extra": "743590 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 1490,
            "unit": "ns/op",
            "extra": "743590 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 660,
            "unit": "B/op",
            "extra": "743590 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "743590 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 15805,
            "unit": "ns/op\t    7327 B/op\t     161 allocs/op",
            "extra": "75135 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 15805,
            "unit": "ns/op",
            "extra": "75135 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 7327,
            "unit": "B/op",
            "extra": "75135 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 161,
            "unit": "allocs/op",
            "extra": "75135 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4439,
            "unit": "ns/op\t    1488 B/op\t      41 allocs/op",
            "extra": "261564 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4439,
            "unit": "ns/op",
            "extra": "261564 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1488,
            "unit": "B/op",
            "extra": "261564 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "261564 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 16962,
            "unit": "ns/op\t   17078 B/op\t     228 allocs/op",
            "extra": "63394 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 16962,
            "unit": "ns/op",
            "extra": "63394 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 17078,
            "unit": "B/op",
            "extra": "63394 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 228,
            "unit": "allocs/op",
            "extra": "63394 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 8441,
            "unit": "ns/op\t    3312 B/op\t     138 allocs/op",
            "extra": "139899 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 8441,
            "unit": "ns/op",
            "extra": "139899 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3312,
            "unit": "B/op",
            "extra": "139899 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "139899 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 8991,
            "unit": "ns/op\t    5776 B/op\t      54 allocs/op",
            "extra": "130615 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 8991,
            "unit": "ns/op",
            "extra": "130615 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5776,
            "unit": "B/op",
            "extra": "130615 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "130615 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1040,
            "unit": "ns/op\t     672 B/op\t      25 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1040,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - B/op",
            "value": 672,
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
            "value": 1537,
            "unit": "ns/op\t     776 B/op\t      31 allocs/op",
            "extra": "690456 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1537,
            "unit": "ns/op",
            "extra": "690456 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 776,
            "unit": "B/op",
            "extra": "690456 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "690456 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1622,
            "unit": "ns/op\t     816 B/op\t      32 allocs/op",
            "extra": "670822 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1622,
            "unit": "ns/op",
            "extra": "670822 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 816,
            "unit": "B/op",
            "extra": "670822 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "670822 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1787,
            "unit": "ns/op\t     944 B/op\t      36 allocs/op",
            "extra": "629103 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1787,
            "unit": "ns/op",
            "extra": "629103 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 944,
            "unit": "B/op",
            "extra": "629103 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "629103 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 2891,
            "unit": "ns/op\t    1512 B/op\t      63 allocs/op",
            "extra": "385958 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 2891,
            "unit": "ns/op",
            "extra": "385958 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1512,
            "unit": "B/op",
            "extra": "385958 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "385958 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5298,
            "unit": "ns/op\t    2824 B/op\t     118 allocs/op",
            "extra": "220314 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5298,
            "unit": "ns/op",
            "extra": "220314 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 2824,
            "unit": "B/op",
            "extra": "220314 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "220314 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3358,
            "unit": "ns/op\t    1896 B/op\t      75 allocs/op",
            "extra": "336537 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3358,
            "unit": "ns/op",
            "extra": "336537 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 1896,
            "unit": "B/op",
            "extra": "336537 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "336537 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1123,
            "unit": "ns/op\t     568 B/op\t      23 allocs/op",
            "extra": "921252 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1123,
            "unit": "ns/op",
            "extra": "921252 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "921252 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "921252 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2251,
            "unit": "ns/op\t    1400 B/op\t      44 allocs/op",
            "extra": "510579 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2251,
            "unit": "ns/op",
            "extra": "510579 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1400,
            "unit": "B/op",
            "extra": "510579 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "510579 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2327,
            "unit": "ns/op\t    1512 B/op\t      46 allocs/op",
            "extra": "479070 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2327,
            "unit": "ns/op",
            "extra": "479070 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1512,
            "unit": "B/op",
            "extra": "479070 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "479070 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1300,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "818948 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1300,
            "unit": "ns/op",
            "extra": "818948 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "818948 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "818948 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1300,
            "unit": "ns/op\t     720 B/op\t      24 allocs/op",
            "extra": "815258 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1300,
            "unit": "ns/op",
            "extra": "815258 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 720,
            "unit": "B/op",
            "extra": "815258 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "815258 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1584,
            "unit": "ns/op\t     848 B/op\t      30 allocs/op",
            "extra": "692252 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1584,
            "unit": "ns/op",
            "extra": "692252 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 848,
            "unit": "B/op",
            "extra": "692252 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "692252 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 13454,
            "unit": "ns/op\t    8640 B/op\t     105 allocs/op",
            "extra": "88934 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 13454,
            "unit": "ns/op",
            "extra": "88934 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 8640,
            "unit": "B/op",
            "extra": "88934 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 105,
            "unit": "allocs/op",
            "extra": "88934 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 36986,
            "unit": "ns/op\t    7449 B/op\t      99 allocs/op",
            "extra": "32287 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 36986,
            "unit": "ns/op",
            "extra": "32287 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7449,
            "unit": "B/op",
            "extra": "32287 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "32287 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 37295,
            "unit": "ns/op\t    7609 B/op\t     108 allocs/op",
            "extra": "31897 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 37295,
            "unit": "ns/op",
            "extra": "31897 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7609,
            "unit": "B/op",
            "extra": "31897 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "31897 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 37071,
            "unit": "ns/op\t    7305 B/op\t     103 allocs/op",
            "extra": "31864 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 37071,
            "unit": "ns/op",
            "extra": "31864 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7305,
            "unit": "B/op",
            "extra": "31864 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "31864 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 16505,
            "unit": "ns/op\t    8193 B/op\t     217 allocs/op",
            "extra": "72962 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 16505,
            "unit": "ns/op",
            "extra": "72962 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8193,
            "unit": "B/op",
            "extra": "72962 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "72962 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 161961605,
            "unit": "ns/op\t334344582 B/op\t  281382 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 161961605,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334344582,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281382,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 46665,
            "unit": "ns/op\t   27497 B/op\t     614 allocs/op",
            "extra": "25452 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 46665,
            "unit": "ns/op",
            "extra": "25452 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 27497,
            "unit": "B/op",
            "extra": "25452 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "25452 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime",
            "value": 6627,
            "unit": "ns/op\t    3857 B/op\t      79 allocs/op",
            "extra": "180942 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - ns/op",
            "value": 6627,
            "unit": "ns/op",
            "extra": "180942 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - B/op",
            "value": 3857,
            "unit": "B/op",
            "extra": "180942 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - allocs/op",
            "value": 79,
            "unit": "allocs/op",
            "extra": "180942 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone",
            "value": 208963,
            "unit": "ns/op\t  337299 B/op\t     224 allocs/op",
            "extra": "5194 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - ns/op",
            "value": 208963,
            "unit": "ns/op",
            "extra": "5194 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - B/op",
            "value": 337299,
            "unit": "B/op",
            "extra": "5194 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "5194 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha",
            "value": 3773,
            "unit": "ns/op\t    1504 B/op\t      42 allocs/op",
            "extra": "303219 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - ns/op",
            "value": 3773,
            "unit": "ns/op",
            "extra": "303219 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - B/op",
            "value": 1504,
            "unit": "B/op",
            "extra": "303219 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - allocs/op",
            "value": 42,
            "unit": "allocs/op",
            "extra": "303219 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric",
            "value": 5275,
            "unit": "ns/op\t    2080 B/op\t      58 allocs/op",
            "extra": "221568 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - ns/op",
            "value": 5275,
            "unit": "ns/op",
            "extra": "221568 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - B/op",
            "value": 2080,
            "unit": "B/op",
            "extra": "221568 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "221568 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode",
            "value": 6001,
            "unit": "ns/op\t    2016 B/op\t      56 allocs/op",
            "extra": "191214 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - ns/op",
            "value": 6001,
            "unit": "ns/op",
            "extra": "191214 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - B/op",
            "value": 2016,
            "unit": "B/op",
            "extra": "191214 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - allocs/op",
            "value": 56,
            "unit": "allocs/op",
            "extra": "191214 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode",
            "value": 7679,
            "unit": "ns/op\t    2641 B/op\t      73 allocs/op",
            "extra": "143198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - ns/op",
            "value": 7679,
            "unit": "ns/op",
            "extra": "143198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - B/op",
            "value": 2641,
            "unit": "B/op",
            "extra": "143198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - allocs/op",
            "value": 73,
            "unit": "allocs/op",
            "extra": "143198 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4172,
            "unit": "ns/op\t    1958 B/op\t      58 allocs/op",
            "extra": "275179 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4172,
            "unit": "ns/op",
            "extra": "275179 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 1958,
            "unit": "B/op",
            "extra": "275179 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "275179 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1303,
            "unit": "ns/op\t     512 B/op\t      16 allocs/op",
            "extra": "813117 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1303,
            "unit": "ns/op",
            "extra": "813117 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 512,
            "unit": "B/op",
            "extra": "813117 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "813117 times\n4 procs"
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
          "id": "c558c3dc78bbcd76f7110c8f6b7b8d0539bf399b",
          "message": "chore: Exclude validator comparison module from Renovate (#62)",
          "timestamp": "2024-12-22T19:11:41+01:00",
          "tree_id": "69b03cf41407c4434784e7ddef65558e97f97881",
          "url": "https://github.com/nobl9/govy/commit/c558c3dc78bbcd76f7110c8f6b7b8d0539bf399b"
        },
        "date": 1736768653178,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 702.6,
            "unit": "ns/op\t     208 B/op\t       6 allocs/op",
            "extra": "1713902 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 702.6,
            "unit": "ns/op",
            "extra": "1713902 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 208,
            "unit": "B/op",
            "extra": "1713902 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1713902 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 809.6,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1494854 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 809.6,
            "unit": "ns/op",
            "extra": "1494854 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1494854 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1494854 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 869.4,
            "unit": "ns/op\t     368 B/op\t      10 allocs/op",
            "extra": "1315640 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 869.4,
            "unit": "ns/op",
            "extra": "1315640 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1315640 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1315640 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 766.5,
            "unit": "ns/op\t     352 B/op\t       8 allocs/op",
            "extra": "1577162 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 766.5,
            "unit": "ns/op",
            "extra": "1577162 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 352,
            "unit": "B/op",
            "extra": "1577162 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1577162 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 834.7,
            "unit": "ns/op\t     344 B/op\t      10 allocs/op",
            "extra": "1415121 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 834.7,
            "unit": "ns/op",
            "extra": "1415121 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 344,
            "unit": "B/op",
            "extra": "1415121 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1415121 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 768.5,
            "unit": "ns/op\t     352 B/op\t       8 allocs/op",
            "extra": "1559139 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 768.5,
            "unit": "ns/op",
            "extra": "1559139 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 352,
            "unit": "B/op",
            "extra": "1559139 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1559139 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties",
            "value": 6693,
            "unit": "ns/op\t    2672 B/op\t      83 allocs/op",
            "extra": "172447 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - ns/op",
            "value": 6693,
            "unit": "ns/op",
            "extra": "172447 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - B/op",
            "value": 2672,
            "unit": "B/op",
            "extra": "172447 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - allocs/op",
            "value": 83,
            "unit": "allocs/op",
            "extra": "172447 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1247,
            "unit": "ns/op\t     488 B/op\t      18 allocs/op",
            "extra": "863923 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1247,
            "unit": "ns/op",
            "extra": "863923 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 488,
            "unit": "B/op",
            "extra": "863923 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "863923 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 171.3,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "6991344 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 171.3,
            "unit": "ns/op",
            "extra": "6991344 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "6991344 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6991344 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1343,
            "unit": "ns/op\t     608 B/op\t      16 allocs/op",
            "extra": "808681 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1343,
            "unit": "ns/op",
            "extra": "808681 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 608,
            "unit": "B/op",
            "extra": "808681 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "808681 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength",
            "value": 1067,
            "unit": "ns/op\t     448 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - ns/op",
            "value": 1067,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - B/op",
            "value": 448,
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
            "value": 1094,
            "unit": "ns/op\t     448 B/op\t      12 allocs/op",
            "extra": "984711 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1094,
            "unit": "ns/op",
            "extra": "984711 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "984711 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "984711 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength",
            "value": 1308,
            "unit": "ns/op\t     608 B/op\t      16 allocs/op",
            "extra": "817450 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1308,
            "unit": "ns/op",
            "extra": "817450 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 608,
            "unit": "B/op",
            "extra": "817450 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "817450 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 1049,
            "unit": "ns/op\t     512 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 1049,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - B/op",
            "value": 512,
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
            "value": 1043,
            "unit": "ns/op\t     512 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - ns/op",
            "value": 1043,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - B/op",
            "value": 512,
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
            "value": 1126,
            "unit": "ns/op\t     528 B/op\t      14 allocs/op",
            "extra": "964480 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1126,
            "unit": "ns/op",
            "extra": "964480 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - B/op",
            "value": 528,
            "unit": "B/op",
            "extra": "964480 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "964480 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength",
            "value": 1047,
            "unit": "ns/op\t     512 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - ns/op",
            "value": 1047,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - B/op",
            "value": 512,
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
            "value": 1048,
            "unit": "ns/op\t     512 B/op\t      12 allocs/op",
            "extra": "994722 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - ns/op",
            "value": 1048,
            "unit": "ns/op",
            "extra": "994722 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - B/op",
            "value": 512,
            "unit": "B/op",
            "extra": "994722 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "994722 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf",
            "value": 1056,
            "unit": "ns/op\t     520 B/op\t      22 allocs/op",
            "extra": "982929 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1056,
            "unit": "ns/op",
            "extra": "982929 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "982929 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "982929 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 7298,
            "unit": "ns/op\t    3088 B/op\t      98 allocs/op",
            "extra": "161464 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 7298,
            "unit": "ns/op",
            "extra": "161464 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3088,
            "unit": "B/op",
            "extra": "161464 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "161464 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2670,
            "unit": "ns/op\t    1048 B/op\t      32 allocs/op",
            "extra": "428857 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2670,
            "unit": "ns/op",
            "extra": "428857 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "428857 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "428857 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1054,
            "unit": "ns/op\t     608 B/op\t      23 allocs/op",
            "extra": "1134657 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1054,
            "unit": "ns/op",
            "extra": "1134657 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - B/op",
            "value": 608,
            "unit": "B/op",
            "extra": "1134657 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "1134657 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty",
            "value": 177.4,
            "unit": "ns/op\t     112 B/op\t       4 allocs/op",
            "extra": "6734018 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 177.4,
            "unit": "ns/op",
            "extra": "6734018 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "6734018 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6734018 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 1511,
            "unit": "ns/op\t     612 B/op\t      20 allocs/op",
            "extra": "705120 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 1511,
            "unit": "ns/op",
            "extra": "705120 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 612,
            "unit": "B/op",
            "extra": "705120 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "705120 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 1519,
            "unit": "ns/op\t     660 B/op\t      20 allocs/op",
            "extra": "731575 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 1519,
            "unit": "ns/op",
            "extra": "731575 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 660,
            "unit": "B/op",
            "extra": "731575 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "731575 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 16002,
            "unit": "ns/op\t    7329 B/op\t     161 allocs/op",
            "extra": "75333 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 16002,
            "unit": "ns/op",
            "extra": "75333 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 7329,
            "unit": "B/op",
            "extra": "75333 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 161,
            "unit": "allocs/op",
            "extra": "75333 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4442,
            "unit": "ns/op\t    1488 B/op\t      41 allocs/op",
            "extra": "270520 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4442,
            "unit": "ns/op",
            "extra": "270520 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1488,
            "unit": "B/op",
            "extra": "270520 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "270520 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 16891,
            "unit": "ns/op\t   17078 B/op\t     228 allocs/op",
            "extra": "71690 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 16891,
            "unit": "ns/op",
            "extra": "71690 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 17078,
            "unit": "B/op",
            "extra": "71690 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 228,
            "unit": "allocs/op",
            "extra": "71690 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 8425,
            "unit": "ns/op\t    3312 B/op\t     138 allocs/op",
            "extra": "140107 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 8425,
            "unit": "ns/op",
            "extra": "140107 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3312,
            "unit": "B/op",
            "extra": "140107 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "140107 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 8999,
            "unit": "ns/op\t    5776 B/op\t      54 allocs/op",
            "extra": "131804 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 8999,
            "unit": "ns/op",
            "extra": "131804 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5776,
            "unit": "B/op",
            "extra": "131804 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "131804 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1047,
            "unit": "ns/op\t     672 B/op\t      25 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1047,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - B/op",
            "value": 672,
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
            "value": 1552,
            "unit": "ns/op\t     776 B/op\t      31 allocs/op",
            "extra": "704643 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1552,
            "unit": "ns/op",
            "extra": "704643 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 776,
            "unit": "B/op",
            "extra": "704643 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "704643 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1622,
            "unit": "ns/op\t     816 B/op\t      32 allocs/op",
            "extra": "668443 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1622,
            "unit": "ns/op",
            "extra": "668443 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 816,
            "unit": "B/op",
            "extra": "668443 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "668443 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1800,
            "unit": "ns/op\t     944 B/op\t      36 allocs/op",
            "extra": "635192 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1800,
            "unit": "ns/op",
            "extra": "635192 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 944,
            "unit": "B/op",
            "extra": "635192 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "635192 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 2912,
            "unit": "ns/op\t    1512 B/op\t      63 allocs/op",
            "extra": "389289 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 2912,
            "unit": "ns/op",
            "extra": "389289 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1512,
            "unit": "B/op",
            "extra": "389289 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "389289 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5340,
            "unit": "ns/op\t    2824 B/op\t     118 allocs/op",
            "extra": "212557 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5340,
            "unit": "ns/op",
            "extra": "212557 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 2824,
            "unit": "B/op",
            "extra": "212557 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "212557 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3396,
            "unit": "ns/op\t    1896 B/op\t      75 allocs/op",
            "extra": "343368 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3396,
            "unit": "ns/op",
            "extra": "343368 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 1896,
            "unit": "B/op",
            "extra": "343368 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "343368 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1134,
            "unit": "ns/op\t     568 B/op\t      23 allocs/op",
            "extra": "964044 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1134,
            "unit": "ns/op",
            "extra": "964044 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "964044 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "964044 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2250,
            "unit": "ns/op\t    1400 B/op\t      44 allocs/op",
            "extra": "508207 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2250,
            "unit": "ns/op",
            "extra": "508207 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1400,
            "unit": "B/op",
            "extra": "508207 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "508207 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2361,
            "unit": "ns/op\t    1512 B/op\t      46 allocs/op",
            "extra": "480991 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2361,
            "unit": "ns/op",
            "extra": "480991 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1512,
            "unit": "B/op",
            "extra": "480991 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "480991 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1318,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "803312 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1318,
            "unit": "ns/op",
            "extra": "803312 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "803312 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "803312 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1304,
            "unit": "ns/op\t     720 B/op\t      24 allocs/op",
            "extra": "838324 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1304,
            "unit": "ns/op",
            "extra": "838324 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 720,
            "unit": "B/op",
            "extra": "838324 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "838324 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1611,
            "unit": "ns/op\t     848 B/op\t      30 allocs/op",
            "extra": "692535 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1611,
            "unit": "ns/op",
            "extra": "692535 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 848,
            "unit": "B/op",
            "extra": "692535 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "692535 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 13533,
            "unit": "ns/op\t    8640 B/op\t     105 allocs/op",
            "extra": "86983 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 13533,
            "unit": "ns/op",
            "extra": "86983 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 8640,
            "unit": "B/op",
            "extra": "86983 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 105,
            "unit": "allocs/op",
            "extra": "86983 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 37096,
            "unit": "ns/op\t    7497 B/op\t      99 allocs/op",
            "extra": "32212 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 37096,
            "unit": "ns/op",
            "extra": "32212 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7497,
            "unit": "B/op",
            "extra": "32212 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "32212 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 37260,
            "unit": "ns/op\t    7609 B/op\t     108 allocs/op",
            "extra": "31914 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 37260,
            "unit": "ns/op",
            "extra": "31914 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7609,
            "unit": "B/op",
            "extra": "31914 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "31914 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 37011,
            "unit": "ns/op\t    7305 B/op\t     103 allocs/op",
            "extra": "32250 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 37011,
            "unit": "ns/op",
            "extra": "32250 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7305,
            "unit": "B/op",
            "extra": "32250 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "32250 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 16405,
            "unit": "ns/op\t    8193 B/op\t     217 allocs/op",
            "extra": "73222 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 16405,
            "unit": "ns/op",
            "extra": "73222 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8193,
            "unit": "B/op",
            "extra": "73222 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "73222 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 171032799,
            "unit": "ns/op\t334344347 B/op\t  281381 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 171032799,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334344347,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281381,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 48583,
            "unit": "ns/op\t   27497 B/op\t     614 allocs/op",
            "extra": "24900 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 48583,
            "unit": "ns/op",
            "extra": "24900 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 27497,
            "unit": "B/op",
            "extra": "24900 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "24900 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime",
            "value": 6614,
            "unit": "ns/op\t    3857 B/op\t      79 allocs/op",
            "extra": "175389 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - ns/op",
            "value": 6614,
            "unit": "ns/op",
            "extra": "175389 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - B/op",
            "value": 3857,
            "unit": "B/op",
            "extra": "175389 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - allocs/op",
            "value": 79,
            "unit": "allocs/op",
            "extra": "175389 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone",
            "value": 224499,
            "unit": "ns/op\t  337299 B/op\t     224 allocs/op",
            "extra": "5196 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - ns/op",
            "value": 224499,
            "unit": "ns/op",
            "extra": "5196 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - B/op",
            "value": 337299,
            "unit": "B/op",
            "extra": "5196 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "5196 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha",
            "value": 3857,
            "unit": "ns/op\t    1504 B/op\t      42 allocs/op",
            "extra": "301120 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - ns/op",
            "value": 3857,
            "unit": "ns/op",
            "extra": "301120 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - B/op",
            "value": 1504,
            "unit": "B/op",
            "extra": "301120 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - allocs/op",
            "value": 42,
            "unit": "allocs/op",
            "extra": "301120 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric",
            "value": 5419,
            "unit": "ns/op\t    2080 B/op\t      58 allocs/op",
            "extra": "217737 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - ns/op",
            "value": 5419,
            "unit": "ns/op",
            "extra": "217737 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - B/op",
            "value": 2080,
            "unit": "B/op",
            "extra": "217737 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "217737 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode",
            "value": 6130,
            "unit": "ns/op\t    2016 B/op\t      56 allocs/op",
            "extra": "193276 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - ns/op",
            "value": 6130,
            "unit": "ns/op",
            "extra": "193276 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - B/op",
            "value": 2016,
            "unit": "B/op",
            "extra": "193276 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - allocs/op",
            "value": 56,
            "unit": "allocs/op",
            "extra": "193276 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode",
            "value": 7818,
            "unit": "ns/op\t    2641 B/op\t      73 allocs/op",
            "extra": "152140 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - ns/op",
            "value": 7818,
            "unit": "ns/op",
            "extra": "152140 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - B/op",
            "value": 2641,
            "unit": "B/op",
            "extra": "152140 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - allocs/op",
            "value": 73,
            "unit": "allocs/op",
            "extra": "152140 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4289,
            "unit": "ns/op\t    1958 B/op\t      58 allocs/op",
            "extra": "271488 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4289,
            "unit": "ns/op",
            "extra": "271488 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 1958,
            "unit": "B/op",
            "extra": "271488 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "271488 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1351,
            "unit": "ns/op\t     512 B/op\t      16 allocs/op",
            "extra": "811350 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1351,
            "unit": "ns/op",
            "extra": "811350 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 512,
            "unit": "B/op",
            "extra": "811350 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "811350 times\n4 procs"
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
          "id": "116e75332ea2767b3828503fc516c34d59399678",
          "message": "feat: Add WithExamples function to rules (#59)\n\n## Motivation\r\n\r\nCurrently some validation rules allow providing examples which are then\r\ndisplayed in the produced error. It would be beneficial to allow\r\nproviding such examples to ANY of the `govy.Rule`.\r\n\r\n## Release Notes\r\n\r\nAdded `WithExamples` method to `govy.Rule` which allows adding examples\r\nto the produced error message.\r\nThe examples are converted to their string representation and added\r\nbetween error message and details, like so: `<message> (e.g.\r\n<examples>); <details>`.\r\n\r\n## Breaking Changes\r\n\r\nThe following rules no longer support providing examples as arguments\r\ndirectly to the rule constructor: `StringMatchRegexp`,\r\n`StringDenyRegexp`, `StringDateTime`. Instead, provide these examples by\r\ncalling `WithExamples` method of the `govy.Rule`, e.g.:\r\n`rules.StringMatchRegexp(regexp.MustCompile(\"^John|Jack$\")).WithExamples(\"John\",\r\n\"Jack\")`.",
          "timestamp": "2024-12-22T19:56:49+01:00",
          "tree_id": "e9240a6e530a64fba5c93523e1e8096425599d7e",
          "url": "https://github.com/nobl9/govy/commit/116e75332ea2767b3828503fc516c34d59399678"
        },
        "date": 1736769199576,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 724.8,
            "unit": "ns/op\t     208 B/op\t       6 allocs/op",
            "extra": "1510057 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 724.8,
            "unit": "ns/op",
            "extra": "1510057 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 208,
            "unit": "B/op",
            "extra": "1510057 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1510057 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 834.3,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1442864 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 834.3,
            "unit": "ns/op",
            "extra": "1442864 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1442864 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1442864 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 902,
            "unit": "ns/op\t     368 B/op\t      10 allocs/op",
            "extra": "1400124 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 902,
            "unit": "ns/op",
            "extra": "1400124 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1400124 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1400124 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 777.9,
            "unit": "ns/op\t     352 B/op\t       8 allocs/op",
            "extra": "1572102 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 777.9,
            "unit": "ns/op",
            "extra": "1572102 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 352,
            "unit": "B/op",
            "extra": "1572102 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1572102 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 847.5,
            "unit": "ns/op\t     344 B/op\t      10 allocs/op",
            "extra": "1414333 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 847.5,
            "unit": "ns/op",
            "extra": "1414333 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 344,
            "unit": "B/op",
            "extra": "1414333 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1414333 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 762.1,
            "unit": "ns/op\t     352 B/op\t       8 allocs/op",
            "extra": "1569691 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 762.1,
            "unit": "ns/op",
            "extra": "1569691 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 352,
            "unit": "B/op",
            "extra": "1569691 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1569691 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties",
            "value": 6787,
            "unit": "ns/op\t    2672 B/op\t      83 allocs/op",
            "extra": "171825 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - ns/op",
            "value": 6787,
            "unit": "ns/op",
            "extra": "171825 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - B/op",
            "value": 2672,
            "unit": "B/op",
            "extra": "171825 times\n4 procs"
          },
          {
            "name": "BenchmarkEqualProperties - allocs/op",
            "value": 83,
            "unit": "allocs/op",
            "extra": "171825 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1261,
            "unit": "ns/op\t     488 B/op\t      18 allocs/op",
            "extra": "876436 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1261,
            "unit": "ns/op",
            "extra": "876436 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 488,
            "unit": "B/op",
            "extra": "876436 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "876436 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 168.6,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "7092213 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 168.6,
            "unit": "ns/op",
            "extra": "7092213 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "7092213 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "7092213 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1381,
            "unit": "ns/op\t     608 B/op\t      16 allocs/op",
            "extra": "814504 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1381,
            "unit": "ns/op",
            "extra": "814504 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 608,
            "unit": "B/op",
            "extra": "814504 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "814504 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength",
            "value": 1116,
            "unit": "ns/op\t     448 B/op\t      12 allocs/op",
            "extra": "908185 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - ns/op",
            "value": 1116,
            "unit": "ns/op",
            "extra": "908185 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "908185 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "908185 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength",
            "value": 1116,
            "unit": "ns/op\t     448 B/op\t      12 allocs/op",
            "extra": "984436 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1116,
            "unit": "ns/op",
            "extra": "984436 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "984436 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "984436 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength",
            "value": 1352,
            "unit": "ns/op\t     608 B/op\t      16 allocs/op",
            "extra": "824804 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1352,
            "unit": "ns/op",
            "extra": "824804 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 608,
            "unit": "B/op",
            "extra": "824804 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "824804 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 1085,
            "unit": "ns/op\t     512 B/op\t      12 allocs/op",
            "extra": "994852 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 1085,
            "unit": "ns/op",
            "extra": "994852 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - B/op",
            "value": 512,
            "unit": "B/op",
            "extra": "994852 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "994852 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength",
            "value": 1090,
            "unit": "ns/op\t     512 B/op\t      12 allocs/op",
            "extra": "993108 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - ns/op",
            "value": 1090,
            "unit": "ns/op",
            "extra": "993108 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - B/op",
            "value": 512,
            "unit": "B/op",
            "extra": "993108 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "993108 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength",
            "value": 1145,
            "unit": "ns/op\t     528 B/op\t      14 allocs/op",
            "extra": "951457 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1145,
            "unit": "ns/op",
            "extra": "951457 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - B/op",
            "value": 528,
            "unit": "B/op",
            "extra": "951457 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "951457 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength",
            "value": 1065,
            "unit": "ns/op\t     512 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - ns/op",
            "value": 1065,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - B/op",
            "value": 512,
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
            "value": 1083,
            "unit": "ns/op\t     512 B/op\t      12 allocs/op",
            "extra": "975964 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - ns/op",
            "value": 1083,
            "unit": "ns/op",
            "extra": "975964 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - B/op",
            "value": 512,
            "unit": "B/op",
            "extra": "975964 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "975964 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf",
            "value": 1154,
            "unit": "ns/op\t     520 B/op\t      22 allocs/op",
            "extra": "950664 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1154,
            "unit": "ns/op",
            "extra": "950664 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "950664 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "950664 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 7532,
            "unit": "ns/op\t    3088 B/op\t      98 allocs/op",
            "extra": "158293 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 7532,
            "unit": "ns/op",
            "extra": "158293 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3088,
            "unit": "B/op",
            "extra": "158293 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "158293 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2727,
            "unit": "ns/op\t    1048 B/op\t      32 allocs/op",
            "extra": "424828 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2727,
            "unit": "ns/op",
            "extra": "424828 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "424828 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "424828 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1037,
            "unit": "ns/op\t     608 B/op\t      23 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1037,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - B/op",
            "value": 608,
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
            "value": 182.7,
            "unit": "ns/op\t     112 B/op\t       4 allocs/op",
            "extra": "6456064 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 182.7,
            "unit": "ns/op",
            "extra": "6456064 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 112,
            "unit": "B/op",
            "extra": "6456064 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6456064 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 669.2,
            "unit": "ns/op\t     257 B/op\t       8 allocs/op",
            "extra": "1791006 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 669.2,
            "unit": "ns/op",
            "extra": "1791006 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 257,
            "unit": "B/op",
            "extra": "1791006 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1791006 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 692.1,
            "unit": "ns/op\t     290 B/op\t       8 allocs/op",
            "extra": "1747862 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 692.1,
            "unit": "ns/op",
            "extra": "1747862 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 290,
            "unit": "B/op",
            "extra": "1747862 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1747862 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 16528,
            "unit": "ns/op\t    6992 B/op\t     182 allocs/op",
            "extra": "72224 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 16528,
            "unit": "ns/op",
            "extra": "72224 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 6992,
            "unit": "B/op",
            "extra": "72224 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 182,
            "unit": "allocs/op",
            "extra": "72224 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4369,
            "unit": "ns/op\t    1488 B/op\t      41 allocs/op",
            "extra": "263050 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4369,
            "unit": "ns/op",
            "extra": "263050 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1488,
            "unit": "B/op",
            "extra": "263050 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "263050 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 20792,
            "unit": "ns/op\t   17862 B/op\t     293 allocs/op",
            "extra": "57117 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 20792,
            "unit": "ns/op",
            "extra": "57117 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 17862,
            "unit": "B/op",
            "extra": "57117 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 293,
            "unit": "allocs/op",
            "extra": "57117 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 8717,
            "unit": "ns/op\t    3312 B/op\t     138 allocs/op",
            "extra": "135272 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 8717,
            "unit": "ns/op",
            "extra": "135272 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3312,
            "unit": "B/op",
            "extra": "135272 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "135272 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 9530,
            "unit": "ns/op\t    5776 B/op\t      54 allocs/op",
            "extra": "123723 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 9530,
            "unit": "ns/op",
            "extra": "123723 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5776,
            "unit": "B/op",
            "extra": "123723 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "123723 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1063,
            "unit": "ns/op\t     672 B/op\t      25 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1063,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - B/op",
            "value": 672,
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
            "value": 1568,
            "unit": "ns/op\t     776 B/op\t      31 allocs/op",
            "extra": "700126 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1568,
            "unit": "ns/op",
            "extra": "700126 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 776,
            "unit": "B/op",
            "extra": "700126 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "700126 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1653,
            "unit": "ns/op\t     816 B/op\t      32 allocs/op",
            "extra": "667886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1653,
            "unit": "ns/op",
            "extra": "667886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 816,
            "unit": "B/op",
            "extra": "667886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "667886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1837,
            "unit": "ns/op\t     944 B/op\t      36 allocs/op",
            "extra": "599601 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1837,
            "unit": "ns/op",
            "extra": "599601 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 944,
            "unit": "B/op",
            "extra": "599601 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "599601 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 2918,
            "unit": "ns/op\t    1512 B/op\t      63 allocs/op",
            "extra": "388954 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 2918,
            "unit": "ns/op",
            "extra": "388954 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1512,
            "unit": "B/op",
            "extra": "388954 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "388954 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5345,
            "unit": "ns/op\t    2824 B/op\t     118 allocs/op",
            "extra": "218358 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5345,
            "unit": "ns/op",
            "extra": "218358 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 2824,
            "unit": "B/op",
            "extra": "218358 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "218358 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3431,
            "unit": "ns/op\t    1896 B/op\t      75 allocs/op",
            "extra": "331936 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3431,
            "unit": "ns/op",
            "extra": "331936 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 1896,
            "unit": "B/op",
            "extra": "331936 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "331936 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1161,
            "unit": "ns/op\t     568 B/op\t      23 allocs/op",
            "extra": "951145 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1161,
            "unit": "ns/op",
            "extra": "951145 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "951145 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "951145 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2329,
            "unit": "ns/op\t    1400 B/op\t      44 allocs/op",
            "extra": "495981 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2329,
            "unit": "ns/op",
            "extra": "495981 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1400,
            "unit": "B/op",
            "extra": "495981 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "495981 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2407,
            "unit": "ns/op\t    1512 B/op\t      46 allocs/op",
            "extra": "467521 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2407,
            "unit": "ns/op",
            "extra": "467521 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1512,
            "unit": "B/op",
            "extra": "467521 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "467521 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1354,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "784208 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1354,
            "unit": "ns/op",
            "extra": "784208 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "784208 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "784208 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1361,
            "unit": "ns/op\t     720 B/op\t      24 allocs/op",
            "extra": "811653 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1361,
            "unit": "ns/op",
            "extra": "811653 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 720,
            "unit": "B/op",
            "extra": "811653 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "811653 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1671,
            "unit": "ns/op\t     848 B/op\t      30 allocs/op",
            "extra": "653928 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1671,
            "unit": "ns/op",
            "extra": "653928 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 848,
            "unit": "B/op",
            "extra": "653928 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "653928 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 13201,
            "unit": "ns/op\t    8640 B/op\t     105 allocs/op",
            "extra": "90417 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 13201,
            "unit": "ns/op",
            "extra": "90417 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 8640,
            "unit": "B/op",
            "extra": "90417 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 105,
            "unit": "allocs/op",
            "extra": "90417 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 37399,
            "unit": "ns/op\t    7497 B/op\t      99 allocs/op",
            "extra": "31806 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 37399,
            "unit": "ns/op",
            "extra": "31806 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7497,
            "unit": "B/op",
            "extra": "31806 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "31806 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 37938,
            "unit": "ns/op\t    7609 B/op\t     108 allocs/op",
            "extra": "31526 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 37938,
            "unit": "ns/op",
            "extra": "31526 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7609,
            "unit": "B/op",
            "extra": "31526 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "31526 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 37288,
            "unit": "ns/op\t    7305 B/op\t     103 allocs/op",
            "extra": "31797 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 37288,
            "unit": "ns/op",
            "extra": "31797 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7305,
            "unit": "B/op",
            "extra": "31797 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "31797 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 17796,
            "unit": "ns/op\t    8193 B/op\t     217 allocs/op",
            "extra": "67596 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 17796,
            "unit": "ns/op",
            "extra": "67596 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8193,
            "unit": "B/op",
            "extra": "67596 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "67596 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 168370907,
            "unit": "ns/op\t334342062 B/op\t  281365 allocs/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 168370907,
            "unit": "ns/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334342062,
            "unit": "B/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281365,
            "unit": "allocs/op",
            "extra": "6 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 46634,
            "unit": "ns/op\t   27497 B/op\t     614 allocs/op",
            "extra": "25437 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 46634,
            "unit": "ns/op",
            "extra": "25437 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 27497,
            "unit": "B/op",
            "extra": "25437 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "25437 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime",
            "value": 6218,
            "unit": "ns/op\t    3632 B/op\t      72 allocs/op",
            "extra": "192205 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - ns/op",
            "value": 6218,
            "unit": "ns/op",
            "extra": "192205 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - B/op",
            "value": 3632,
            "unit": "B/op",
            "extra": "192205 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDateTime - allocs/op",
            "value": 72,
            "unit": "allocs/op",
            "extra": "192205 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone",
            "value": 210071,
            "unit": "ns/op\t  338219 B/op\t     280 allocs/op",
            "extra": "5444 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - ns/op",
            "value": 210071,
            "unit": "ns/op",
            "extra": "5444 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - B/op",
            "value": 338219,
            "unit": "B/op",
            "extra": "5444 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTimeZone - allocs/op",
            "value": 280,
            "unit": "allocs/op",
            "extra": "5444 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha",
            "value": 3698,
            "unit": "ns/op\t    1504 B/op\t      42 allocs/op",
            "extra": "309978 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - ns/op",
            "value": 3698,
            "unit": "ns/op",
            "extra": "309978 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - B/op",
            "value": 1504,
            "unit": "B/op",
            "extra": "309978 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlpha - allocs/op",
            "value": 42,
            "unit": "allocs/op",
            "extra": "309978 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric",
            "value": 5152,
            "unit": "ns/op\t    2080 B/op\t      58 allocs/op",
            "extra": "229366 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - ns/op",
            "value": 5152,
            "unit": "ns/op",
            "extra": "229366 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - B/op",
            "value": 2080,
            "unit": "B/op",
            "extra": "229366 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumeric - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "229366 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode",
            "value": 5861,
            "unit": "ns/op\t    2016 B/op\t      56 allocs/op",
            "extra": "199886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - ns/op",
            "value": 5861,
            "unit": "ns/op",
            "extra": "199886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - B/op",
            "value": 2016,
            "unit": "B/op",
            "extra": "199886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphaUnicode - allocs/op",
            "value": 56,
            "unit": "allocs/op",
            "extra": "199886 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode",
            "value": 7565,
            "unit": "ns/op\t    2641 B/op\t      73 allocs/op",
            "extra": "155420 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - ns/op",
            "value": 7565,
            "unit": "ns/op",
            "extra": "155420 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - B/op",
            "value": 2641,
            "unit": "B/op",
            "extra": "155420 times\n4 procs"
          },
          {
            "name": "BenchmarkStringAlphanumericUnicode - allocs/op",
            "value": 73,
            "unit": "allocs/op",
            "extra": "155420 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4301,
            "unit": "ns/op\t    1958 B/op\t      58 allocs/op",
            "extra": "263016 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4301,
            "unit": "ns/op",
            "extra": "263016 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 1958,
            "unit": "B/op",
            "extra": "263016 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "263016 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1319,
            "unit": "ns/op\t     512 B/op\t      16 allocs/op",
            "extra": "829860 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1319,
            "unit": "ns/op",
            "extra": "829860 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 512,
            "unit": "B/op",
            "extra": "829860 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "829860 times\n4 procs"
          }
        ]
      }
    ]
  }
}