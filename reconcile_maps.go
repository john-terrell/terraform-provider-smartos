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
		if oldKeyIndex >= len(oldKeys) {
			break
		}

		oldKey := oldKeys[oldKeyIndex]

		// Get the new key at the same index.
		log.Printf("COMPARE: oldKeyIndex '%d' oldKey '%s'\n", oldKeyIndex, oldKey)

		// if there are no  more old keys, the remaining new keys need to be added
		if oldKeyIndex >= len(oldKeys) {
			break
		}

		// If there are no more new keys, the remaining old keys need to be removed.
		if newKeyIndex >= len(newKeys) {
			break
		}

		newKey := newKeys[newKeyIndex]

		log.Printf("COMPARING old key '%s' with new key '%s'\n", oldKey, newKey)

		if oldKey == newKey {
			// Keys are the same, see if the values are the same
			oldValue := oldMap[oldKey]
			newValue := newMap[oldKey]

			if !itemsAreEqual(oldValue, newValue) {
				// Value changed
				updateItem(oldKey, newValue)
				log.Printf("COMPARE: Updating customer metadata [%s] = %s", oldKey, newValue.(string))
				changesMade = true
			} else {
				log.Printf("COMPARE: No change to customer metadata [%s] = %s", oldKey, newValue.(string))
			}
			newKeyIndex++
			oldKeyIndex++
		} else if oldKey < newKey {
			// 'oldKey' was removed
			oldValue := oldMap[oldKey].(string)
			removeItem(oldKey)
			log.Printf("COMPARE: Removing customer metadata [%s] = %s", oldKey, oldValue)
			changesMade = true

			oldKeyIndex++
		} else {
			// 'newKey' was added
			newValue := newMap[newKey]

			addItem(newKey, newValue)
			log.Printf("COMPARE: Adding customer metadata [%s] = %s", newKey, newValue.(string))

			newKeyIndex++
			changesMade = true
		}
	}

	// Remove any remaining old keys
	for ; oldKeyIndex < len(oldKeys); oldKeyIndex++ {
		oldKey := oldKeys[oldKeyIndex]
		oldValue := oldMap[oldKey].(string)
		removeItem(oldKey)
		log.Printf("COMPARE: Removing customer metadata at end [%s] = %s", oldKey, oldValue)
		changesMade = true
	}

	// Add any remaining new keys
	for ; newKeyIndex < len(newKeys); newKeyIndex++ {
		newKey := newKeys[newKeyIndex]
		newValue := newMap[newKey]

		addItem(newKey, newValue)
		log.Printf("COMPARE: Adding customer metadata at end [%s] = %s", newKey, newValue.(string))
		changesMade = true
	}

	return changesMade
}
