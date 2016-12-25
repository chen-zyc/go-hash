package hash

import (
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

var testKey = "test-key"

func _TestHashHit(t *testing.T) {
	keys := randomKeys(100, 16)
	for _, f := range AllFuncs {
		t.Log("")
		t.Log(getFuncName(f))
		hit := make([]int, 10)
		for _, key := range keys {
			hit[int(f(key)%uint64(len(hit)))]++
		}
		barChart(t, hit)
	}
}

func getFuncName(f interface{}) string {
	longName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	parts := strings.Split(longName, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return "?"
}

func randomKeys(n, m int) (keys []string) {
	rand.Seed(time.Now().Unix())
	keys = make([]string, n)
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < n; i++ {
		key := make([]byte, m)
		for j := 0; j < m; j++ {
			key[j] = s[rand.Intn(len(s))]
		}
		keys[i] = string(key)
	}
	return
}

func repeat(s string, i int) string {
	v := ""
	for k := 0; k < i; k++ {
		v += s
	}
	return v
}

func barChart(t *testing.T, hit []int) {
	for _, n := range hit {
		t.Log(repeat("0", n))
	}
}

func BenchmarkRSHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RSHash(testKey)
	}
}

func BenchmarkJSHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JSHash(testKey)
	}
}

func BenchmarkPJWHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PJWHash(testKey)
	}
}

func BenchmarkELFHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ELFHash(testKey)
	}
}

func BenchmarkBKDRHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BKDRHash(testKey)
	}
}

func BenchmarkSDBMHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SDBMHash(testKey)
	}
}

func BenchmarkDJBHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DJBHash(testKey)
	}
}

func BenchmarkDEKHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DEKHash(testKey)
	}
}

func BenchmarkAPHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		APHash(testKey)
	}
}

func BenchmarkFNVHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FNVHash(testKey)
	}
}

func BenchmarkMYSQLHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MYSQLHash(testKey)
	}
}
