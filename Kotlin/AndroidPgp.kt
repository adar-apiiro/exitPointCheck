import com.github.kibotu.android_pgp.AndroidPgp

fun main() {
    val message = "Hello, PGP!"
    val publicKey = "YourPublicKeyHere" // Replace with an actual PGP public key
    val privateKey = "YourPrivateKeyHere" // Replace with an actual PGP private key
    val passphrase = "YourPassphrase"

    // Encrypt
    val encrypted = AndroidPgp.encrypt(message, publicKey)
    println("Encrypted: $encrypted")

    // Decrypt
    val decrypted = AndroidPgp.decrypt(encrypted, privateKey, passphrase)
    println("Decrypted: $decrypted")
}
