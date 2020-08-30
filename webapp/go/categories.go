package main

// +----+-----------+--------------------------------+
// | id | parent_id | category_name                  |
// +----+-----------+--------------------------------+
// |  1 |         0 | ソファー                       |
// |  2 |         1 | 一人掛けソファー               |
// |  3 |         1 | 二人掛けソファー               |
// |  4 |         1 | コーナーソファー               |
// |  5 |         1 | 二段ソファー                   |
// |  6 |         1 | ソファーベッド                 |
// | 10 |         0 | 家庭用チェア                   |
// | 11 |        10 | スツール                       |
// | 12 |        10 | クッションスツール             |
// | 13 |        10 | ダイニングチェア               |
// | 14 |        10 | リビングチェア                 |
// | 15 |        10 | カウンターチェア               |
// | 20 |         0 | キッズチェア                   |
// | 21 |        20 | 学習チェア                     |
// | 22 |        20 | ベビーソファ                   |
// | 23 |        20 | キッズハイチェア               |
// | 24 |        20 | テーブルチェア                 |
// | 30 |         0 | オフィスチェア                 |
// | 31 |        30 | デスクチェア                   |
// | 32 |        30 | ビジネスチェア                 |
// | 33 |        30 | 回転チェア                     |
// | 34 |        30 | リクライニングチェア           |
// | 35 |        30 | 投擲用椅子                     |
// | 40 |         0 | 折りたたみ椅子                 |
// | 41 |        40 | パイプ椅子                     |
// | 42 |        40 | 木製折りたたみ椅子             |
// | 43 |        40 | キッチンチェア                 |
// | 44 |        40 | アウトドアチェア               |
// | 45 |        40 | 作業椅子                       |
// | 50 |         0 | ベンチ                         |
// | 51 |        50 | 一人掛けベンチ                 |
// | 52 |        50 | 二人掛けベンチ                 |
// | 53 |        50 | アウトドア用ベンチ             |
// | 54 |        50 | 収納付きベンチ                 |
// | 55 |        50 | 背もたれ付きベンチ             |
// | 56 |        50 | ベンチマーク                   |
// | 60 |         0 | 座椅子                         |
// | 61 |        60 | 和風座椅子                     |
// | 62 |        60 | 高座椅子                       |
// | 63 |        60 | ゲーミング座椅子               |
// | 64 |        60 | ロッキングチェア               |
// | 65 |        60 | 座布団                         |
// | 66 |        60 | 空気椅子                       |
// +----+-----------+--------------------------------+
var CategoriesMap map[int]Category
var CategoriesIDsPerParent map[int][]int

func init() {
	CategoriesIDsPerParent = map[int][]int{
		0:  []int{1, 10, 20, 30, 40, 50, 60},
		1:  []int{2, 3, 4, 5, 6},
		10: []int{11, 12, 13, 14},
		20: []int{21, 22, 23, 24},
		30: []int{31, 32, 33, 34, 35},
		40: []int{41, 42, 43, 44, 45},
		50: []int{51, 52, 53, 54, 55, 56},
		60: []int{61, 62, 63, 64, 65, 66},
	}

	CategoriesMap = map[int]Category{
		1: {
			ID:                 1,
			ParentID:           0,
			CategoryName:       "ソファー",
			ParentCategoryName: "",
		},
		2: {
			ID:                 2,
			ParentID:           1,
			CategoryName:       "一人掛けソファー",
			ParentCategoryName: "ソファー",
		},
		3: {
			ID:                 3,
			ParentID:           1,
			CategoryName:       "二人掛けソファー",
			ParentCategoryName: "ソファー",
		},
		4: {
			ID:                 4,
			ParentID:           1,
			CategoryName:       "コーナーソファー",
			ParentCategoryName: "ソファー",
		},
		5: {
			ID:                 5,
			ParentID:           1,
			CategoryName:       "二段ソファー",
			ParentCategoryName: "ソファー",
		},
		6: {
			ID:                 6,
			ParentID:           1,
			CategoryName:       "ソファーベッド",
			ParentCategoryName: "ソファー",
		},

		10: {
			ID:                 10,
			ParentID:           0,
			CategoryName:       "家庭用チェア",
			ParentCategoryName: "",
		},
		11: {
			ID:                 11,
			ParentID:           10,
			CategoryName:       "スツール",
			ParentCategoryName: "家庭用チェア",
		},
		12: {
			ID:                 12,
			ParentID:           10,
			CategoryName:       "クッションスツール",
			ParentCategoryName: "家庭用チェア",
		},
		13: {
			ID:                 13,
			ParentID:           10,
			CategoryName:       "ダイニングチェア",
			ParentCategoryName: "家庭用チェア",
		},
		14: {
			ID:                 14,
			ParentID:           10,
			CategoryName:       "リビングチェア",
			ParentCategoryName: "家庭用チェア",
		},
		15: {
			ID:                 15,
			ParentID:           10,
			CategoryName:       "カウンターチェア",
			ParentCategoryName: "家庭用チェア",
		},

		20: {
			ID:                 20,
			ParentID:           0,
			CategoryName:       "キッズチェア",
			ParentCategoryName: "",
		},
		21: {
			ID:                 21,
			ParentID:           20,
			CategoryName:       "学習チェア",
			ParentCategoryName: "キッズチェア",
		},
		22: {
			ID:                 22,
			ParentID:           20,
			CategoryName:       "ベビーソファ",
			ParentCategoryName: "キッズチェア",
		},
		23: {
			ID:                 23,
			ParentID:           20,
			CategoryName:       "キッズハイチェア",
			ParentCategoryName: "キッズチェア",
		},
		24: {
			ID:                 24,
			ParentID:           20,
			CategoryName:       "テーブルチェア",
			ParentCategoryName: "キッズチェア",
		},

		30: {
			ID:                 30,
			ParentID:           0,
			CategoryName:       "オフィスチェア",
			ParentCategoryName: "",
		},
		31: {
			ID:                 31,
			ParentID:           30,
			CategoryName:       "デスクチェア",
			ParentCategoryName: "オフィスチェア",
		},
		32: {
			ID:                 32,
			ParentID:           30,
			CategoryName:       "ビジネスチェア",
			ParentCategoryName: "オフィスチェア",
		},
		33: {
			ID:                 33,
			ParentID:           30,
			CategoryName:       "回転チェア",
			ParentCategoryName: "オフィスチェア",
		},
		34: {
			ID:                 34,
			ParentID:           30,
			CategoryName:       "リクライニングチェア",
			ParentCategoryName: "オフィスチェア",
		},
		35: {
			ID:                 35,
			ParentID:           30,
			CategoryName:       "投擲用椅子",
			ParentCategoryName: "オフィスチェア",
		},

		40: {
			ID:                 40,
			ParentID:           0,
			CategoryName:       "折りたたみ椅子",
			ParentCategoryName: "",
		},
		41: {
			ID:                 41,
			ParentID:           40,
			CategoryName:       "パイプ椅子",
			ParentCategoryName: "折りたたみ椅子",
		},
		42: {
			ID:                 42,
			ParentID:           40,
			CategoryName:       "木製折りたたみ椅子",
			ParentCategoryName: "折りたたみ椅子",
		},
		43: {
			ID:                 43,
			ParentID:           40,
			CategoryName:       "キッチンチェア",
			ParentCategoryName: "折りたたみ椅子",
		},
		44: {
			ID:                 44,
			ParentID:           40,
			CategoryName:       "アウトドアチェア",
			ParentCategoryName: "折りたたみ椅子",
		},
		45: {
			ID:                 45,
			ParentID:           40,
			CategoryName:       "作業椅子",
			ParentCategoryName: "折りたたみ椅子",
		},

		50: {
			ID:                 50,
			ParentID:           0,
			CategoryName:       "ベンチ",
			ParentCategoryName: "",
		},
		51: {
			ID:                 51,
			ParentID:           50,
			CategoryName:       "一人掛けベンチ",
			ParentCategoryName: "ベンチ",
		},
		52: {
			ID:                 52,
			ParentID:           50,
			CategoryName:       "二人掛けベンチ",
			ParentCategoryName: "ベンチ",
		},
		53: {
			ID:                 53,
			ParentID:           50,
			CategoryName:       "アウトドア用ベンチ",
			ParentCategoryName: "ベンチ",
		},
		54: {
			ID:                 54,
			ParentID:           50,
			CategoryName:       "収納付きベンチ",
			ParentCategoryName: "ベンチ",
		},
		55: {
			ID:                 55,
			ParentID:           50,
			CategoryName:       "背もたれ付きベンチ",
			ParentCategoryName: "ベンチ",
		},
		56: {
			ID:                 56,
			ParentID:           50,
			CategoryName:       "ベンチマーク",
			ParentCategoryName: "ベンチ",
		},

		60: {
			ID:                 60,
			ParentID:           0,
			CategoryName:       "座椅子",
			ParentCategoryName: "",
		},
		61: {
			ID:                 61,
			ParentID:           60,
			CategoryName:       "和風座椅子",
			ParentCategoryName: "座椅子",
		},
		62: {
			ID:                 62,
			ParentID:           60,
			CategoryName:       "高座椅子",
			ParentCategoryName: "座椅子",
		},
		63: {
			ID:                 63,
			ParentID:           60,
			CategoryName:       "ゲーミング座椅子",
			ParentCategoryName: "座椅子",
		},
		64: {
			ID:                 64,
			ParentID:           60,
			CategoryName:       "ロッキングチェア",
			ParentCategoryName: "座椅子",
		},
		65: {
			ID:                 65,
			ParentID:           60,
			CategoryName:       "座布団",
			ParentCategoryName: "座椅子",
		},
		66: {
			ID:                 66,
			ParentID:           60,
			CategoryName:       "空気椅子",
			ParentCategoryName: "座椅子",
		},
	}
}
