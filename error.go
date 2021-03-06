package witch

import "errors"

var (
    UnknownType   = errors.New("No needed types in cases!\n")
    UnknownMethod = errors.New("Magic Piu change the value failed, cause of the unknown given method!\n")
    UnAddrAndSet  = errors.New("Unaddressable and unsettable value!\n")
    UnAddressable = errors.New("Unaddressable value!\n")
    UnSettable    = errors.New("Unsettable value!\n")
    UnmatchedType = errors.New("Unmatched type!\n")
)
