/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file gonlpir_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Mar  8 15:40:16 2015
 *
 **/

package gonlpir

import (
	"fmt"
	"testing"
)

//===================================================================
// Public APIs
//===================================================================

func TestParagraphProcess(t *testing.T) {
	nlpir, err := NewNLPIR(UTF8, "")
	if err != nil {
		t.Error(err)
		return
	}
	defer nlpir.Exit()
	fmt.Println("TestParagraphProcess: 我是 ３中国人")
	fmt.Println(nlpir.ParagraphProcess("我是 ３中国人", true))
}

func TestParagraphProcessA(t *testing.T) {
	nlpir, err := NewNLPIR(UTF8, "")
	if err != nil {
		t.Error(err)
		return
	}
	defer nlpir.Exit()
	fmt.Println("TestParagraphProcessA: == 字幕制作：波士顿 == 你听我的劝，吃点好吗？　")
	results := nlpir.ParagraphProcessA("== 字幕制作：波士顿 == 你听我的劝，吃点好吗？　", true)
	for i := 0; i < len(results); i++ {
		fmt.Printf("%#v\n", results[i])
	}
}

//===================================================================
// Private
//===================================================================
