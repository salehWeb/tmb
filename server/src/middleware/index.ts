import { Request, Response, NextFunction } from "express"
import jwt from "jsonwebtoken";

export const Authorized = (req: Request, res: Response, next: NextFunction) => {
    try {
        if (req.cookies['token']) {
            jwt.verify(req.cookies['token'], process.env.SECRET_KEY as string, (err: any, decodedToken: any) => {
                if (err) return res.status(401).json({ massage: "Unauthorized" })

                return next({ id: Number(decodedToken.id) })
            });
        }
        return res.status(401).json({ massage: "Unauthorized" })
    } catch (error) {
        console.log(error)
        return res.status(500).json({ massage: "internal Server Error" })
    }
}
