/*
 * Inflector Pkg (Go)
 *
 * Copyright (c) 2013 Ivan Torres
 * Released under the MIT license
 * https://github.com/mexpolk/inflector/blob/master/LICENSE
 *
 */

package inflector

import (
	"testing"
)

type inflectionSample struct {
	str, out string
}

func TestToCamel(t *testing.T) {
	samples := []inflectionSample{
		{"sample text", "sampleText"},
		{"sample-text", "sampleText"},
		{"sample_text", "sampleText"},
		{"sampleText", "sampleText"},
		{"sample 2 Text", "sample2Text"},
	}

	for _, sample := range samples {
		if out := ToCamel(sample.str); out != sample.out {
			t.Errorf("got %q, expected %q", out, sample.out)
		}
	}
}

func TestToDash(t *testing.T) {
	samples := []inflectionSample{
		{"sample text", "sample-text"},
		{"sample-text", "sample-text"},
		{"sample_text", "sample-text"},
		{"sampleText", "sample-text"},
		{"sample 2 Text", "sample-2-text"},
	}

	for _, sample := range samples {
		if out := ToDash(sample.str); out != sample.out {
			t.Errorf("got %q, expected %q", out, sample.out)
		}
	}
}

func TestToPascal(t *testing.T) {
	samples := []inflectionSample{
		{"sample text", "SampleText"},
		{"sample-text", "SampleText"},
		{"sample_text", "SampleText"},
		{"sampleText", "SampleText"},
		{"sample 2 Text", "Sample2Text"},
	}

	for _, sample := range samples {
		if out := ToPascal(sample.str); out != sample.out {
			t.Errorf("got %q, expected %q", out, sample.out)
		}
	}
}

func TestToUnderscore(t *testing.T) {
	samples := []inflectionSample{
		{"sample text", "sample_text"},
		{"sample-text", "sample_text"},
		{"sample_text", "sample_text"},
		{"sampleText", "sample_text"},
		{"sample 2 Text", "sample_2_text"},
		{"SAMPLE 2 TEXT", "sample_2_text"},
		{"Base64Encode", "base64_encode"},
		{"FOO:BAR$BAZ", "foo_bar_baz"},
		{"something.com", "something_com"},
	}

	for _, sample := range samples {
		if out := ToUnderscore(sample.str); out != sample.out {
			t.Errorf("got %q, expected %q", out, sample.out)
		}
	}
}
