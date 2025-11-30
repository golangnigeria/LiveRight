// utils/uploadToCloudinary.js
export async function UploadToCloudinary(file) {
  const formData = new FormData();
  formData.append('file', file);
  formData.append('upload_preset', 'mynnefarm'); // set in Cloudinary dashboard

  const res = await fetch('https://api.cloudinary.com/v1_1/dvmoamnui/image/upload', {
    method: 'POST',
    body: formData,
  });

  const data = await res.json();
  console.log('Cloudinary response:', data);
  return data.secure_url;
}
