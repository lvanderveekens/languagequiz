import { NextApiRequest, NextApiResponse } from "next";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Exercise[] | ErrorResponse>
) {
  try {
    const userId = req.query.userId;
    const response = await fetch(`http://localhost:8888/v1/exercises`);
    const data = await response.json();
    res.status(200).json(data);
  } catch (err) {
    console.error(err);
    res.status(500).json({ message: "Internal Server Error" });
  }
}

type Exercise = { };

type ErrorResponse = {
  message: string
}