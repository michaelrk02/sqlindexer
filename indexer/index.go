package indexer

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

type Index struct {
	Name      string
	GroupName string
	Table     string
	Fields    []string
}

func GetKeyName(table, field string) string {
	h := md5.New()
	_, _ = h.Write([]byte(fmt.Sprintf("%s%s", table, field)))

	return fmt.Sprintf(
		"ix_%s_%s_%s",
		shrink(table),
		shrink(field),
		string(hex.EncodeToString(h.Sum(nil)))[:6],
	)
}

func (idx *Index) ID() string {
	return strings.Join(idx.Fields, "|")
}

func (idx *Index) SQL() string {
	fields := make([]string, len(idx.Fields))
	for i, field := range idx.Fields {
		fields[i] = "`" + field + "`"
	}

	return fmt.Sprintf(
		"ALTER TABLE `%s` ADD INDEX `%s` (%s)",
		idx.Table,
		idx.Name,
		strings.Join(fields, ", "),
	)
}

func shrink(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		parts[i] = part[:min(3, len(part))]
	}
	return strings.Join(parts, "")
}
