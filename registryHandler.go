package main

import (
	"fmt"

	clog "github.com/charmbracelet/log"
	"golang.org/x/sys/windows/registry"
)

func checkRegistry() {
	_, err := registry.OpenKey(registry.CURRENT_USER, "Software\\Classes\\SystemFileAssociations\\.png\\shell\\Kpixel", registry.QUERY_VALUE)
	if err != nil {
		createRegistry()
	}
}

func createRegistry() {
	// clog.Info("couldn't access the registry, creating one...")
	pngReg, _, err := registry.CreateKey(registry.CURRENT_USER, "Software\\Classes\\SystemFileAssociations\\.png\\shell\\Kpixel", registry.ALL_ACCESS)
	if err != nil {
		clog.Fatal(err)
	}
	defer pngReg.Close()

	pngReg.SetStringValue("MUIVerb", "Kpixel")
	pngReg.SetStringValue("Icon", "Kpixel.ico")
	pngReg.SetStringValue("subCommands", "")

	ColumnSortReg, _, err := registry.CreateKey(pngReg, "Software\\Classes\\SystemFileAssociations\\.png\\shell\\Kpixel\\ColumnSort", registry.ALL_ACCESS)
	if err != nil {
		clog.Fatal(err)
	}
	defer ColumnSortReg.Close()

	ColumnSortReg.SetStringValue("MUIVerb", "Column Sort")

	ColumnSortCommandReg, _, err := registry.CreateKey(pngReg, "Software\\Classes\\SystemFileAssociations\\.png\\shell\\Kpixel\\ColumnSort\\command", registry.ALL_ACCESS)
	if err != nil {
		clog.Fatal(err)
	}
	defer ColumnSortCommandReg.Close()

	ColumnSortCommandReg.SetStringValue("@", "cmd /k \"\"D:\\resources\\maketx.exe\" \"%1\"\"")

	clog.Info(pngReg.Stat())
}

func cleanRegistry(registryKey string) error {
	var access uint32 = registry.QUERY_VALUE | registry.ENUMERATE_SUB_KEYS
	regKey, err := registry.OpenKey(registry.LOCAL_MACHINE, registryKey, access)
	if err != nil {
		if err == registry.ErrNotExist {
			return nil
		}

		fmt.Printf("Failed to open '%s' Error: %v", registryKey, err)
		return err
	}

	defer func() {
		if err := regKey.Close(); err != nil {
			fmt.Printf("Failed to close reg key '%v' Error: %v", regKey, err)
		}
	}()

	keyNames, err := regKey.ReadSubKeyNames(0)
	if err != nil {
		fmt.Printf("Failed to get %q keys from registry error: %v", regKey, err)
		return nil
	}

	if err := deleteSubKeys(keyNames, registryKey, access); err != nil {
		fmt.Printf("delete error: %v", err)
		return err
	}

	// delete itself
	if err := registry.DeleteKey(regKey, ""); err != nil {
		fmt.Printf("Cannot delete key path : %s Error: %v", registryKey, err)
		return err
	}
	return nil
}

func deleteSubKeys(keyNames []string, registryKey string, access uint32) error {
	for _, k := range keyNames {
		keyPath := fmt.Sprintf("%s\\%s", registryKey, k)

		subRegKey, err := registry.OpenKey(registry.LOCAL_MACHINE, keyPath, access)
		if err != nil {
			fmt.Printf("Path %q not found on registry: %v", keyPath, err)
			return err
		}

		// delete sub keys
		if err := deleteSubKeysRecursively(subRegKey, keyPath, access); err != nil {
			fmt.Printf("deleteSubKeysRecursively error: %v", err)
		}

		if err := registry.DeleteKey(subRegKey, ""); err != nil {
			fmt.Printf("Cannot delete key path : %s Error: %v", keyPath, err)
			return err
		}
	}

	return nil
}

func deleteSubKeysRecursively(regKey registry.Key, keyPath string, access uint32) error {
	subKeyNames, err := regKey.ReadSubKeyNames(0)
	if err != nil {
		return nil
	}

	for _, subKeyName := range subKeyNames {
		path := fmt.Sprintf("%s\\%s", keyPath, subKeyName)

		subRegKey, err := registry.OpenKey(registry.LOCAL_MACHINE, path, access)
		if err != nil {
			fmt.Printf("Path %q not found on registry: %v", subKeyName, err)
			return err
		}

		if err = deleteSubKeysRecursively(subRegKey, path, access); err != nil {
			return err
		}

		if err = registry.DeleteKey(subRegKey, ""); err != nil {
			fmt.Printf("Cannot delete registry key path : %s Error: %v", path, err)
			return err
		}
	}

	return nil
}
