package src

import (
	"log"
	"os/exec"
	"strconv"
	"time"
)

func Logic(waktu TimeShutdown) {
	active := true
	log.Printf("PC Akan mati pada tanggal %d jam %d menit %d", waktu.Tanggal, waktu.Jam, waktu.Menit)
	for active {
		tn := GetTimeNow()
		if waktu.Tanggal <= tn.Tanggal && waktu.Jam <= tn.Jam && waktu.Menit <= tn.Menit {
			log.Print("times out")
			active = false
			cmd := exec.Command("cmd", "/C", "shutdown", "/s")
			if err := cmd.Run(); err != nil {
				log.Print(err)
			}
		}
		time.Sleep(60 * time.Second)
	}
}

func ParamToTime(waktu []string) TimeShutdown {
	valueInt := []int{}

	if len(waktu) == 0 {
		dt := time.Now()
		besok := dt.Day() + 1
		return TimeShutdown{besok, 1, 0}
	}

	length := len(waktu)

	for i := 0; i < 3; i++ {
		if length <= i {
			valueInt = append(valueInt, 0)
			continue
		}

		value, err := strconv.Atoi(waktu[i])
		if err != nil {
			log.Print(err)
		}

		valueInt = append(valueInt, value)
	}

	return TimeShutdown{valueInt[0], valueInt[1], valueInt[2]}
}

func GetTimeNow() TimeShutdown {
	dt := time.Now()
	return TimeShutdown{
		dt.Day(),
		dt.Hour(),
		dt.Minute(),
	}
}
