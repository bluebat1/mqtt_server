export default {
    post(url:string, map:any, fn:Function, error:Function) {
        let fd = new FormData()
        for (const key in map) {
            fd.append(key, map[key])
        }
        (window as any).axios
            .post(
                url,
                fd,
                {}
            )
            .then((e:any) => {
                let data = e.data;
                fn(data)
            })
            .catch((e:any) => {
                if (e.response || e.request) {
                    error()
                }
            });
    },
    postAwait(url: string, map:any):any  {
        return new Promise((resolve, reject) => {
            let fd = new FormData()
            for (const key in map) {
                fd.append(key, map[key])
            }
            (window as any).axios
                .post(
                    url,
                    fd,
                    {}
                )
                .then((e:any) => {
                    let data = e.data;
                    resolve(data)
                })
                .catch((e:any) => {
                    if (e.response || e.request) {
                        reject(e)
                    }
                });
        });
    }
}