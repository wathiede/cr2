package cr2

// Fetch the sample files via HTTP so we don't have to check them in to git.
// There is about ~280MB of files listed below.
var mirrorHost = "https://googledrive.com/host/0B7S33WfKVjF0dTZhMGxVaUhvcnc/"
var samples = []sample{
	sample{
		File:        "sample.cr2",
		OriginalURL: "http://nf.wh3rd.net/img/sample.cr2",
		MirrorURL:   mirrorHost + "sample.cr2",
		ThumbW:      5184,
		ThumbH:      3456,
		Filesize:    25881839,
		Checksum:    "bfcff428b7700926f03393aff8fe3bc84046d5e3",
	},
	sample{
		File:        "RAW_CANON_G10.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/g10/RAW_CANON_G10.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_G10.CR2",
		RawW:        4416,
		RawH:        3312,
		ThumbW:      1600,
		ThumbH:      1200,
		Filesize:    20092969,
		Checksum:    "69ae7a8d5cee96906bf05f324a11a86209a43aed",
	},
	sample{
		File:        "RAW_CANON_400D_ARGB.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/400d/RAW_CANON_400D_ARGB.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_400D_ARGB.CR2",
		RawW:        3888,
		RawH:        2592,
		ThumbW:      1936,
		ThumbH:      1288,
		Filesize:    10916213,
		Checksum:    "965319c524de701eeaa11d7c0a78692e704169d7",
	},
	sample{
		File:        "RAW_CANON_1DMARK3.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/1dm3/RAW_CANON_1DMARK3.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_1DMARK3.CR2",
		RawW:        3888,
		RawH:        2592,
		ThumbW:      1936,
		ThumbH:      1288,
		Filesize:    14590280,
		Checksum:    "8dd54604e50343f84e3006ae57f7d9988e9ccdb8",
	},
	sample{
		File:        "RAW_CANON_20D.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/20d/RAW_CANON_20D.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_20D.CR2",
		RawW:        3504,
		RawH:        2336,
		ThumbW:      1536,
		ThumbH:      1024,
		Filesize:    7193541,
		Checksum:    "f8fbe4175ddd309e5dd0bb1821dd519c221029b8",
	},
	sample{
		File:        "RAW_CANON_1DM2N.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/1dm2n/RAW_CANON_1DM2N.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_1DM2N.CR2",
		RawW:        3504,
		RawH:        2336,
		ThumbW:      1728,
		ThumbH:      1152,
		Filesize:    7316992,
		Checksum:    "46a50c8d592362cfaddac8640648b29134f20ec9",
	},
	sample{
		File:        "RAW_CANON_30D.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/30d/RAW_CANON_30D.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_30D.CR2",
		RawW:        3504,
		RawH:        2336,
		ThumbW:      1728,
		ThumbH:      1152,
		Filesize:    7699286,
		Checksum:    "f3f3f47eae3e300548c5ac86594ee2db1d70db03",
	},
	sample{
		File:        "RAW_CANON_1DSM2.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/1dsm2/RAW_CANON_1DSM2.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_1DSM2.CR2",
		RawW:        4992,
		RawH:        3328,
		ThumbW:      1536,
		ThumbH:      1024,
		Filesize:    14914910,
		Checksum:    "b86061ea8fac318db674a804886c1977ec07023a",
	},
	sample{
		File:        "RAW_CANON_5D_ARGB.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/5d/RAW_CANON_5D_ARGB.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_5D_ARGB.CR2",
		RawW:        4368,
		RawH:        2912,
		ThumbW:      2496,
		ThumbH:      1664,
		Filesize:    11138246,
		Checksum:    "0fcd2eb8c629899e3b53e1f4c45d344439da4780",
	},
	sample{
		File:        "RAW_CANON_1DM2.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/1dm2/RAW_CANON_1DM2.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_1DM2.CR2",
		RawW:        3504,
		RawH:        2336,
		ThumbW:      1536,
		ThumbH:      1024,
		Filesize:    6953301,
		Checksum:    "22f5589e376e1030d21e63ae00ca09e537092e74",
	},
	sample{
		File:        "RAW_CANON_350D.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/350d/RAW_CANON_350D.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_350D.CR2",
		RawW:        3456,
		RawH:        2304,
		ThumbW:      1536,
		ThumbH:      1024,
		Filesize:    6717063,
		Checksum:    "1d6203e45d3928a0299044ad0b8538be785e76b9",
	},
	sample{
		File:        "RAW_CANON_G9.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/g9/RAW_CANON_G9.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_G9.CR2",
		RawW:        4000,
		RawH:        3000,
		ThumbW:      1600,
		ThumbH:      1200,
		Filesize:    13952063,
		Checksum:    "25a6d4c65932c66cd100233f3fba233d9f77e5e1",
	},
	sample{
		File:        "RAW_CANON_1DSM3.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/1dsm3/RAW_CANON_1DSM3.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_1DSM3.CR2",
		RawW:        5616,
		RawH:        3744,
		ThumbW:      2784,
		ThumbH:      1856,
		Filesize:    20859768,
		Checksum:    "f9b78010aae4db394a6eb42985f9851104ec7810",
	},
	sample{
		File:        "RAW_CANON_450D.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/450d/RAW_CANON_450D.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_450D.CR2",
		RawW:        4272,
		RawH:        2848,
		ThumbW:      2256,
		ThumbH:      1504,
		Filesize:    17710643,
		Checksum:    "85da365a765e0a3161d58c806541e2462e8fd538",
	},
	sample{
		File:        "RAW_CANON_40D_RAW_V105.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/40d/RAW_CANON_40D_RAW_V105.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_40D_RAW_V105.CR2",
		RawW:        3888,
		RawH:        2592,
		ThumbW:      1936,
		ThumbH:      1288,
		Filesize:    12399795,
		Checksum:    "bf9e939aa38cd93014a659dc4dcdc017d37d534f",
	},
	sample{
		File:        "RAW_CANON_40D_SRAW_V103.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/40d/RAW_CANON_40D_SRAW_V103.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_40D_SRAW_V103.CR2",
		RawW:        1936,
		RawH:        1288,
		ThumbW:      1936,
		ThumbH:      1288,
		Filesize:    6800865,
		Checksum:    "a44b1a4476f1cf0f3ca2e20b5b66d07945ececb1",
	},
	sample{
		File:        "RAW_CANON_40D_RAW_V103.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/40d/RAW_CANON_40D_RAW_V103.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_40D_RAW_V103.CR2",
		RawW:        3888,
		RawH:        2592,
		ThumbW:      1936,
		ThumbH:      1288,
		Filesize:    11330256,
		Checksum:    "9c2d4885123f2b5551b1646c0d182ff51bde7ff7",
	},
	sample{
		File:        "RAW_CANON_40D_RAW_V104.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/40d/RAW_CANON_40D_RAW_V104.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_40D_RAW_V104.CR2",
		RawW:        3888,
		RawH:        2592,
		ThumbW:      1936,
		ThumbH:      1288,
		Filesize:    11468332,
		Checksum:    "5429b38df6f8de09d5b379e7dbd1ae4cc26d6664",
	},
	sample{
		File:        "RAW_CANON_40D_RAW_V336643C.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/40d/RAW_CANON_40D_RAW_V336643C.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_40D_RAW_V336643C.CR2",
		RawW:        3888,
		RawH:        2592,
		ThumbW:      1936,
		ThumbH:      1288,
		Filesize:    13227614,
		Checksum:    "30d24fa9a5773745c6a2dc90e4e55c4a858b1abd",
	},
	sample{
		File:        "RAW_CANON_50D.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/50d/RAW_CANON_50D.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_50D.CR2",
		RawW:        4752,
		RawH:        3168,
		ThumbW:      4752,
		ThumbH:      3168,
		Filesize:    18935626,
		Checksum:    "5ec4b5b5dc828a197a70303337838801917a02a1",
	},
	sample{
		File:        "RAW_CANON_5DMARK2_PREPROD.CR2",
		OriginalURL: "http://www.rawsamples.ch/raws/canon/5dm2/RAW_CANON_5DMARK2_PREPROD.CR2",
		MirrorURL:   mirrorHost + "RAW_CANON_5DMARK2_PREPROD.CR2",
		RawW:        5616,
		RawH:        3744,
		ThumbW:      5616,
		ThumbH:      3744,
		Filesize:    26374329,
		Checksum:    "f7a5ea5e0ca970fa4861689db35bf5d266e04fa1",
	},
}
