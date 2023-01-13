import { NextApiRequest, NextApiResponse } from "next"

const handler = async (req: NextApiRequest, res: NextApiResponse) => {
    if (req.method === "POST") {
        res.status(201).json({massage: "Created Chanel", chanelId: "aisbu67fsavtiasyubvssa8"})
    }
}

export default handler;
