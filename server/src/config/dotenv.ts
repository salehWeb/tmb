import fs from 'fs'

export default function LoadEnv() {
    fs.readFile("./.env", (err, buf) => {
        if (err) {
            console.log("No env File Found")
            return
        }

        const data = String(buf).trim().split("\n")

        for (let item of data) {
            const data = item.trim().split("=\"")

            if (data.length < 2) {
                return
            }

            const kay = data[0]
            // use slice to remove this character __"__ in the end of string
            const value = data[1].slice(0, data[1].length - 1)

            process.env[kay] = value
        }
    })
}
