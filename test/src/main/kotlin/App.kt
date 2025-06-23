import com.example.mytest.Test
import com.example.mytest.myMessage
import com.github.mtps.protobuf.crypt.CryptProvider
import com.github.mtps.protobuf.crypt.encrypted
import com.google.protobuf.Any
import com.google.protobuf.StringValue
import java.util.Base64

@OptIn(ExperimentalStdlibApi::class)
fun main() {
    CryptProvider.register(
        { Base64.getEncoder().encode(it) },
        { Base64.getDecoder().decode(it) },
    )

    val any = Any.pack(StringValue.newBuilder().setValue("hello").build())

    val kt = myMessage {
        encryptedBytes = any.toByteArray().encrypted()
        encryptedInt = 123.encrypted()

        encryptedBytes = encrypted(any.toByteArray())
        encryptedInt = encrypted(123)

        encryptedBytes = encrypted { any.toByteArray() }
        encryptedInt = encrypted { 123 }
    }


    123.encrypted()
    val msg = Test.MyMessage.newBuilder()
        .setEncryptedBytes("Test".toByteArray(Charsets.UTF_8))
//        .setEncryptedAny(any)
//        .setEncryptedBytes("test".toByteArray())
//        .setEncryptedBytes("test".toByteArray())
//        .setEncryptedString("test")
//        .setEncryptedTimestamp(Timestamp.newBuilder().setSeconds(123).setNanos(456).build())
        .build()

//    println("enc: " + msg.encryptedInt.value)
//    println("dec: " + msg.decryptEncryptedInt())
//    println()

//    println("enc: " + msg.encryptedAny.value.toByteArray().toHexString())
//    println("dec: " + msg.decryptEncryptedAny())
//    println()

    println("enc: " + msg.encryptedBytes.value.toByteArray().toHexString())
    println("dec: " + msg.encryptedBytes.decrypt().decodeToString())
    println()

//    println("enc: " + msg.encryptedString.value.toByteArray().toHexString())
//    println("dec: " + msg.decryptEncryptedString())
//    println()

//    println("enc: " + msg.encryptedTimestamp.value.toByteArray().toHexString())
//    println("dec: " + msg.decryptEncryptedTimestamp())
//    println()

}
