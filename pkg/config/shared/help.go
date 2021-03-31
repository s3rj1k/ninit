package shared

import (
	"fmt"
	"strings"
)

// Help prints user-friendly help with available configuration options.
func Help(name, version, buildTime, envPrefix, descriptionBoody string) {
	if name == "" {
		name = UnknownValue
	}

	if version == "" {
		version = UnknownValue
	}

	if buildTime == "" {
		buildTime = UnknownValue
	}

	r := strings.NewReplacer(
		"\t", "  ",
		"%PREFIX%", envPrefix,
		"%NAME%", name,
		"%VERSION%", version,
		"%BUILD_TIME%", buildTime,
	)

	descr := strings.TrimPrefix(DescriptionPrefix, "\n") + "\n" + strings.TrimPrefix(descriptionBoody, "\n")
	descr = r.Replace(descr)

	fmt.Println(descr) //nolint: forbidigo // print help to stdout
}
