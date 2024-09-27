package enums

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type ValidatorFunc func(v string) (bool, []string)

var (
	validatorDimension = func(v string) (bool, []string) {
		if strings.HasSuffix(v, "%") {
			percentage, err := fmt.Sscanf(strings.TrimPrefix(v, "%"), "%d", new(int))
			if err != nil {
				return false, nil
			}

			if percentage <= 1 || percentage >= 100 {
				return false, nil
			}

			return true, nil
		}

		if _, err := strconv.Atoi(v); err != nil {
			return false, nil
		}

		return true, nil
	}

	validatorColour = func(v string) (bool, []string) {
		validColors := []string{
			"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white",
			"brightred", "brightgreen", "brightyellow", "default", "terminal",
		}

		// Check if color is in the list of valid named colors
		for _, validColor := range validColors {
			if v == validColor {
				return true, nil
			}
		}

		// Check for "colour0" to "colour255"
		if strings.HasPrefix(v, "colour") {
			numPart := strings.TrimPrefix(v, "colour")
			if n, err := fmt.Sscanf(numPart, "%d", new(int)); err == nil && n >= 0 && n <= 255 {
				return true, nil
			}
		}

		// Check for hexadecimal RGB strings like '#FFFFFF'
		hexColorPattern := regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$`)
		if hexColorPattern.MatchString(v) {
			return true, nil
		}

		return false, nil
	}

	validatorDefault = func(v string) (bool, []string) { return true, nil }

	validatorToggle = func(v string) (bool, []string) {
		return validatorContains(v, "on", "off")
	}

	validatorIsNumber = func(v string) (bool, []string) {
		integer, err := strconv.Atoi(v)
		if err != nil {
			return false, nil
		}

		if integer < 0 {
			return false, nil
		}

		return true, nil
	}

	validatorNonEmpty = func(v string) (bool, []string) {
		if len(strings.TrimSpace(v)) != 1 {
			return false, nil
		}

		return true, nil
	}

	validatorContains = func(v string, o ...string) (bool, []string) {
		return slices.Contains(o, v), o
	}
)
