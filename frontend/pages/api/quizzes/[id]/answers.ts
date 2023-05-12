import { NextApiRequest, NextApiResponse } from "next";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  try {
    const { id } = req.query

    const response = await fetch(`http://localhost:8888/v1/quizzes/${id}/answers`, {
      method: 'POST',
      body: JSON.stringify(req.body),
      headers: {
        'Content-Type': 'application/json'
      }
    });

    if (response.ok) {
      const json = await response.json();
      console.log(json);
      res.status(200).json(json);
    } else {
      const text = await response.text();
      console.error(text);
      res.status(500).json({ message: "Internal Server Error" });
    }
  } catch (err: any) {
    console.error(err);
    res.status(500).json({ message: "Internal Server Error" });
  }
}

type Exercise = { };