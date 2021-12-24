package dapofiles

import (
	"downfiles" //"github.com/vibrill/downfiles"
	"strconv"
	"strings"
)

func Cek() (siswa, guru, tendik string) {
	files := downfiles.DownloadFiles()
	var (
		filelist  []string
		sislist   []string
		gurulist  []string
		tendilist []string
		gurumax   int
		sismax    int
		tendmax   int
	)

	//membuat file list dari file file yang dideteksi sebagai file dapodik
	for _, f := range files {
		//fmt.Println(f.Name())
		if len(f.Name()) > 70 {
			if f.Name()[len(f.Name())-5:] == ".xlsx" {
				if f.Name()[:7] == "daftar_" || f.Name()[:7] == "daftar-" {
					filelist = append(filelist, f.Name())
				}
			}
		}
	}

	if len(filelist) != 0 {
		//memggolongkan file berdasarkan siswa, guru atau tendik
		for _, f := range filelist {
			//fmt.Println(f)

			if f[:9] == "daftar_gu" || f[:9] == "daftar-gu" {
				gurulist = append(gurulist, f)
			}
			if f[:9] == "daftar_pd" || f[:9] == "daftar-pd" {
				sislist = append(sislist, f)
			}
			if f[:9] == "daftar_te" || f[:9] == "daftar-te" {
				tendilist = append(tendilist, f)
			}
		}
		//fmt.Println(gurulist, sislist, tendilist)

		//pilih int tanggal terkini

		if len(gurulist) != 0 {
			var guruint []int
			for _, f := range gurulist {
				guruint = append(guruint, getdate(f))
			}
			gurumax = findMax(guruint)
			for _, f := range gurulist {
				if getdate(f) == gurumax {
					guru = f
				}
			}
		} else {
			guru = "empty"
		}

		if len(sislist) != 0 {
			var sisint []int
			for _, f := range sislist {
				sisint = append(sisint, getdate(f))
			}
			sismax = findMax(sisint)
			for _, f := range sislist {
				if getdate(f) == sismax {
					siswa = f
				}
			}
		} else {
			siswa = "empty"
		}

		if len(tendilist) != 0 {
			var tendint []int
			for _, f := range tendilist {
				tendint = append(tendint, getdate(f))
			}
			tendmax = findMax(tendint)
			for _, f := range tendilist {
				if getdate(f) == tendmax {
					tendik = f
				}
			}
		} else {
			tendik = "empty"
		}
	} else {
		siswa = "empty"
		guru = "empty"
		tendik = "empty"
	}
	return siswa, guru, tendik

}

func getdate(f string) int {
	k := strings.ReplaceAll(f, `_`, ``)
	k = strings.ReplaceAll(k, `-`, ``)
	k = strings.ReplaceAll(k, ` `, ``)
	k = k[len(k)-16 : len(k)-7]
	intK, _ := strconv.Atoi(k)
	return intK
}

func findMax(a []int) int {
	max := 0
	if len(a) != 0 {
		max = a[0]
		for _, i := range a {
			if max < i {
				max = i
			}
		}
	}
	return max
}
