import { NextApiRequest, NextApiResponse } from "next";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  if (req.method === "POST") {
    try {
      const response = await fetch(`http://localhost:8888/v1/quizzes`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(req.body),
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
  } else if (req.method == "GET") {
    try {
      const response = await fetch(`http://localhost:8888/v1/quizzes`);
      const data = await response.json();
      res.status(200).json(data);
    } catch (err) {
      console.error(err);
      res.status(500).json({ message: "Internal Server Error" });
    }
  } else {
    res.status(415).json({ message: "Unsupported Media Type" });
  }
}