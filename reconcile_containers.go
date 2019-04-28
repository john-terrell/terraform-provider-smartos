package main

import (
	"log"
	"sort"
)

type AddItemFunc func(string, interface{})
type UpdateItemFunc func(string, interface{})
type ItemEqualFunc func(interface{}, interface{}) bool
type RemoveItemFunc func(string)

// ReconcileMaps compares two maps; one with old data and one with new data and calls the various argument functions reflecting what operations are necessary to transform the old map into the new.
func ReconcileMaps(oldMap map[string]interface{}, newMap map[string]interface{}, addItem AddItemFunc, updateItem UpdateItemFunc, removeItem RemoveItemFunc, itemsAreEqual ItemEqualFunc) bool {
	changesMade := false

	var oldKeys []string
	var newKeys []string

	for key := range oldMap {
		oldKeys = append(oldKeys, key)
	}
	sort.Strings(oldKeys)

	for key := range newMap {
		newKeys = append(newKeys, key)
	}
	sort.Strings(newKeys)

	newKeyIndex := 0
	oldKeyIndex := 0

	for {
		// if there are no more old keys, the remaining new keys need to be added
		if oldKeyIndex >= len(oldKeys) {
			break
		}

		// If there are no more new keys, the remaining old keys need to be removed.
		if newKeyIndex >= len(newKeys) {
			break
		}

		oldKey := oldKeys[oldKeyIndex]
		newKey := newKeys[newKeyIndex]

		log.Printf("COMPARING old key '%s' with new key '%s'\n", oldKey, newKey)

		if oldKey == newKey {
			// Keys are the same, see if the values are the same
			oldValue := oldMap[oldKey]
			newValue := newMap[oldKey]

			if !itemsAreEqual(oldValue, newValue) {
				// Value changed
				updateItem(oldKey, newValue)
				log.Printf("COMPARE: Updating [%s] = %s", oldKey, newValue.(string))
				changesMade = true
			} else {
				log.Printf("COMPARE: No change to [%s] = %s", oldKey, newValue.(string))
			}
			newKeyIndex++
			oldKeyIndex++
		} else if oldKey < newKey {
			// 'oldKey' was removed
			oldValue := oldMap[oldKey].(string)
			removeItem(oldKey)
			log.Printf("COMPARE: Removing [%s] = %s", oldKey, oldValue)
			changesMade = true

			oldKeyIndex++
		} else {
			// 'newKey' was added
			newValue := newMap[newKey]

			addItem(newKey, newValue)
			log.Printf("COMPARE: Adding [%s] = %s", newKey, newValue.(string))

			newKeyIndex++
			changesMade = true
		}
	}

	// Remove any remaining old keys
	for ; oldKeyIndex < len(oldKeys); oldKeyIndex++ {
		oldKey := oldKeys[oldKeyIndex]
		oldValue := oldMap[oldKey].(string)
		removeItem(oldKey)
		log.Printf("COMPARE: Removing at end [%s] = %s", oldKey, oldValue)
		changesMade = true
	}

	// Add any remaining new keys
	for ; newKeyIndex < len(newKeys); newKeyIndex++ {
		newKey := newKeys[newKeyIndex]
		newValue := newMap[newKey]

		addItem(newKey, newValue)
		log.Printf("COMPARE: Adding at end [%s] = %s", newKey, newValue.(string))
		changesMade = true
	}

	return changesMade
}

type AddSliceItemFunc func(string)
type RemoveSliceItemFunc func(string)

// ReconcileSlices compares two slices; one with old data and one with new data and calls the various argument functions reflecting what operations are necessary to transform the old slice into the new.
func ReconcileSlices(oldSlice []string, newSlice []string, addItem AddSliceItemFunc, removeItem RemoveSliceItemFunc) bool {
	changesMade := false

	sort.Strings(oldSlice)
	sort.Strings(newSlice)

	newSliceIndex := 0
	oldSliceIndex := 0

	for {
		// if there are no more old values, the remaining new values need to be added
		if oldSliceIndex >= len(oldSlice) {
			break
		}

		// If there are no more new values, the remaining old values need to be removed.
		if newSliceIndex >= len(newSlice) {
			break
		}

		oldValue := oldSlice[oldSliceIndex]
		newValue := newSlice[newSliceIndex]

		log.Printf("COMPARING old value '%s' with new value '%s'\n", oldValue, newValue)

		if oldValue == newValue {
			newSliceIndex++
			oldSliceIndex++
		} else if oldValue < newValue {
			// 'oldValue' was removed
			removeItem(oldValue)
			log.Printf("COMPARE: Removing [%s]", oldValue)
			changesMade = true

			oldSliceIndex++
		} else {
			// 'newValue' was added
			addItem(newValue)
			log.Printf("COMPARE: Adding [%s]", newValue)

			newSliceIndex++
			changesMade = true
		}
	}

	// Remove any remaining old values
	for ; oldSliceIndex < len(oldSlice); oldSliceIndex++ {
		oldValue := oldSlice[oldSliceIndex]
		removeItem(oldValue)
		log.Printf("COMPARE: Removing [%s]", oldValue)
		changesMade = true
	}

	// Add any remaining new values
	for ; newSliceIndex < len(newSlice); newSliceIndex++ {
		newValue := newSlice[newSliceIndex]
		addItem(newValue)
		log.Printf("COMPARE: Adding at end [%s]", newValue)
		changesMade = true
	}

	return changesMade
}
