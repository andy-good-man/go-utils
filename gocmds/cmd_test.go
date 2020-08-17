package gocmds

import (
	"encoding/json"
	"fmt"
	"testing"
)

type TestFFprobe struct {
	Streams []*TestStream
	Format  *TestFormat
}

type TestStream struct {
	Index     int    `json:"index"`
	CodecType string `json:"codec_type"`

	// video
	Width         int    `json:"width"`
	Height        int    `json:"height"`
	CodedWidth    int    `json:"coded_width"`
	CodedHeight   int    `json:"coded_height"`
	StartTime     string `json:"start_time"`
	Duration      string `json:"duration"`
	BitRate       string `json:"bit_rate"`
	RealFrameRate string `json:"r_frame_rate"`
	AvgFrameRate  string `json:"avg_frame_rate"`
	TotalFrame    string `json:"nb_frames"`
}

type TestFormat struct {
	FileName   string `json:"filename"`
	StreamNum  int    `json:"nb_streams"`
	FormatName string `json:"format_name"`
	StartTime  string `json:"start_time"`
	Duration   string `json:"duration"`
	Size       string `json:"size"`
	BitRate    string `json:"bit_rate"`
}

func TestRun(t *testing.T) {
	cmder := new(Cmder)

	// ret := cmder.Run("pwd && ls -rtl", 4)
	// ret := cmder.Run("ls -rtl", 4)
	cmdline := "ffprobe -v quiet -print_format json -show_format -show_streams ./test/debug-1.mp4"
	// ret := cmder.Run(cmdline)
	ret := cmder.Run(cmdline, 10)

	// fmt.Println("ret  :", ret)
	fmt.Println("ret err  :", ret.Err()) // nil ==> if ret.Err() == nil {}
	fmt.Println("ret  :", ret)

	ffprobe := new(TestFFprobe)
	retByte := ret.buf.Bytes()
	err := json.Unmarshal(retByte, ffprobe)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("ffprobe  :%+v\n", ffprobe)

	marshal, _ := json.Marshal(ffprobe)
	fmt.Println("marshal  :", string(marshal))

	fmt.Println("ffprobe.Streams[0].Index  :", ffprobe.Streams[0].Index)
	fmt.Println("ffprobe.Streams[0].CodecType  :", ffprobe.Streams[0].CodecType)
	fmt.Println("ffprobe.Streams[0].RealFrameRate  :", ffprobe.Streams[0].RealFrameRate)
	fmt.Println("ffprobe.Format.FileName  :", ffprobe.Format.FileName)
	fmt.Println("ffprobe.Format.Size  :", ffprobe.Format.Size)
	fmt.Println("ffprobe.Format.Duration  :", ffprobe.Format.Duration)
}
