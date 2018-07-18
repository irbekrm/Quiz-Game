package main

import "testing"

var testQuizzItem = QuizzItem{}
var testRecord = []string{"5+5", "10"}

func TestsetValues(t *testing.T){

  (&testQuizzItem).setValues(testRecord)
  if testQuizzItem.question != testRecord[0] || testQuizzItem.answer != testRecord[1] {
    t.Errorf("Expected \"5+5\" and \"10\", got %v and %v\n", testQuizzItem.question, testQuizzItem.answer)
  }
}

func TestisAnswerCorrect(t *testing.T){
  (&testQuizzItem).setValues(testRecord)
  var answer bool
  
  answer = testQuizzItem.isAnswerCorrect("9")
  if answer {
    t.Errorf("Expected false, got %v\n", answer)
  }

  answer = testQuizzItem.isAnswerCorrect("10")
  if !answer {
    t.Errorf("Expected true, got %v\n", answer)
  }
}
