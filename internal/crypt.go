package types

import "strings"

var CryptProviderPkg = "com.github.mtps.protobuf.crypt"
var CryptProviderFileKt = strings.ReplaceAll(CryptProviderPkg, ".", "/") + "/CryptProviderKt.kt"
var CryptProviderFileJava = strings.ReplaceAll(CryptProviderPkg, ".", "/") + "/CryptProvider.java"
var CryptProviderName = CryptProviderPkg + ".CryptProvider"
