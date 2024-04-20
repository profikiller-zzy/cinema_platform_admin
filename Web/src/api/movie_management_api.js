import {Service} from "@/service/request.js";

export function movieCreateApi(data) {
    const formData = new FormData();
    formData.append('movie_name', data.movie_name);
    formData.append('release_date', data.release_date);
    formData.append('play_time', data.play_time);
    formData.append('director', data.director);
    formData.append('actors', data.actors);
    formData.append('coverPicture', data.coverPicture); // Assuming coverPicture is a File object

    return axios.post("/api/movies", formData, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
}

// 编辑用户
export function movieEditApi(data) {
    return Service.put("/api/movies", data)
}