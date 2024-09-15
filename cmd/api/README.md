# CONCEPTS IMPLEMENTED

## A. CONTEXT

- **Context** dalam Go berhubungan erat dengan HTTP request karena fungsinya untuk mengelola data dan status terkait suatu request secara aman dan efisien selama siklus hidup request tersebut. Berikut adalah alasan mengapa context digunakan bersama HTTP request
- Mengelola Data Selama Siklus Hidup Request
- Menyimpan Data Kunci secara Lokal di Request menggunakan `context.WithValue()`
- Propagasi Status dan Pembatalan (Cancellation Propagation) untuk membatalkan Request
- Pembatasan Batas Waktu (Timeout/Deadline) untuk mengatur batas waktu Request sehingga efisien untk kinerja server
- Meningkatkan Modularitas dan Clean Code karena menghindari passing argumen antar fungsi dalam golang

### Kesimpulan

context digunakan dengan HTTP request untuk mengelola data, status, dan sinyal pembatalan yang terkait dengan request tersebut secara terpusat dan aman selama siklus hidup request, sehingga memungkinkan penanganan yang lebih bersih, efisien, dan terstruktur

## B. POINTER

```go
   var numberA int = 4
   var numberB *int = &numberA

   fmt.Println("numberA (value)   :", numberA)  // 4
   fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

   fmt.Println("numberB (value)   :", *numberB) // 4
   fmt.Println("numberB (address) :", numberB)  // 0xc20800a220
```
