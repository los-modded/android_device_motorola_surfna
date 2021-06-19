package surfna

import (
    "android/soong/android"
)

func init() {
    android.RegisterModuleType("motorola_surfna_init_library_static", initLibraryFactory)
}
