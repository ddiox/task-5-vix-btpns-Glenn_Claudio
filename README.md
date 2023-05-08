# task-5-vix-btpns-Glenn_Claudio
Merancang API pada fitur upload, dan menghapus gambar

Berdasarkan data yang telah diolah oleh tim Data Analysts, bahwa untuk
meningkatkan engagement user pada aplikasi m-banking adalah meningkatkan
aspek memiliki user pada aplikasi tersebut. Saran yang diberikan oleh tim data
analysts adalah membentuk fitur personalize user, salah satunya adalah
memungkinkan user dapat mengupload gambar untuk dijadikan foto profilnya. Tim
developer bertanggung jawab untuk mengembangkan fitur ini, dan kalian diberikan
tugas untuk merancang API pada fitur upload, dan menghapus gambar. Beberapa
ketentuannya antara lain :

a. User dapat menambahkan foto profile

b. Sistem dapat mengidentifikasi User ( log in / sign up)

c. Hanya user yang telah login / sign up yang dapat melakukan delete / tambah
foto profil

d. User dapat menghapus gambar yang telah di post

e. User yang berbeda tidak dapat menghapus / mengubah foto yang telah di
buat oleh user lain

- Ketentuan API :

Pada bagian User Endpoint :

    POST : /users/register, dan gunakan atribut berikut ini
    
    ID (primary key, required)
    
    Username (required)
    
    Email (unique & required) 
    
    Password (required & minlength 6)
    
    Relasi dengan model Photo (Gunakan constraint cascade)
    
    Created At (timestamp)
    
    Updated At (timestamp)
    
    GET: /users/login
    
    Using email & password (required
    
    PUT : /users/:userId (Update User)
    
    DELETE : /users/:userId (Delete User)

Photos Endpoint

    POST : /photos 
    
    ID
    
    Title
    
    Caption
    
    PhotoUrl
    
    UserID
    
    Relasi dengan model User
    
    GET : /photos
    
    PUT : /photoId
    
    DELETE : /:photoId



