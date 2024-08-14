# go-kit-template

Bước 1: Tạo hàm muốn expose ra endpoint trong file service.go
Bước 2: Tạo 1 repo implement các hàm cùng 1 chủ đề
Bước 3: Tạo logic tương ứng cho các hàm (logic tương tác vs db hoặc gọi API sang bên khác) trong folder postgres, mongo, ..., Tạo 1 struct A bọc quanh nó
Bước 4: Bọc quanh struct A bằng Repo B tương ứng
Bước 5: Bọc quanh Repo B bằng Service
Bước 6: Định nghĩa dữ liệu request đầu vào trong file request, thêm hàm validate()
Bước 7: Định nghĩa các route ứng với các endpoint, decode, encode tương ứng trong file transport
Bước 8: Định nghĩa logic các hàm decode (lấy dữ liệu từ body, query_params, grpc,... của request), encode response
Bước 9: Bọc quanh các Service bằng các Endpoint tương ứng
Bước 10: Thêm các thông tin khác như CORS, Content-type ,...
