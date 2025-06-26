package utils

import (
	"math"
	"strings"
)

func tokenize(sentence string) []string {
	return strings.Fields(strings.ToLower(sentence))
}

func buildVocabulary(sentences []string) []string {
  vocabSet := map[string]int{}

  for _, sentence := range sentences {
    for _, word := range tokenize(sentence) {
      vocabSet[word] = 0
    }
  }

  vocabulary := []string{}
  for word := range vocabSet {
    vocabulary = append(vocabulary, word)
  }

  return vocabulary
}

func vectorize(setence string, vocabulary []string) []float64 {
  vector := make([]float64, len(vocabulary))
  wordFreq := map[string]int{}

  for _, word := range tokenize(setence) {
    wordFreq[word]++
  }

  for i, vocab := range vocabulary {
    vector[i] = float64(wordFreq[vocab])
  }

  return vector
}

func vectorizeSentences(sentences, vocabulary []string) [][]float64 {
  vectors := [][]float64{}
  for _, sentence := range sentences {
    vector := vectorize(sentence, vocabulary)
    vectors = append(vectors, vector)
  }

  return vectors
}

func dotProduct(vec1, vec2 []float64) float64 {
  var product float64
  for i := range vec1 {
    product += vec1[i] * vec2[i]
  }

  return product
}

func getMagnitude(vec []float64) float64 {
  var sqrMagnitude float64
  for _, num := range vec {
    sqrMagnitude += num * num
  }

  return math.Sqrt(sqrMagnitude)
}

func CosineSimilarity(vec1, vec2 []float64) float64 {
  dot := dotProduct(vec1, vec2)
  vec1Magnitude := getMagnitude(vec1)
  vec2Magnitude := getMagnitude(vec2)
  
  if vec1Magnitude == 0 || vec2Magnitude == 0 {
    return 0
  }

  cosine := dot / (vec1Magnitude * vec2Magnitude)

  return cosine
}

func FindMostSimilarText(input string, searchList []string) (int, float64) {
  allSentences := make([]string, len(searchList))
  copy(allSentences, searchList)
  allSentences = append(allSentences, input)
  vocabulary := buildVocabulary(allSentences)

  vectors := vectorizeSentences(allSentences, vocabulary)
  inputVector := vectors[len(vectors) - 1]
  similarities := []float64{}

  for i := 0; i < len(vectors) - 1; i++ {
    similarities = append(similarities, CosineSimilarity(inputVector, vectors[i]))
  }

  var maxIndex int
  for i := 1; i < len(similarities); i++ {
    if similarities[maxIndex] < similarities[i] {
      maxIndex = i
    }
  }
  return maxIndex, similarities[maxIndex]
}
