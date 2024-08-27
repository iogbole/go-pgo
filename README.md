
> Read the full blog at [https://israelo.io/blog/pgo/](https://israelo.io/blog/pgo/)

## Introduction 

Over the past few months, I've had numerous discussions with practioners and colleagues on the benefits of Profile-Guided Optimization (PGO). While it’s a topic that generates significant interest, many find it challenging to get started or simply lack the time to explore it fully. As a **Product Manager** in the continuous profiling domain, my curiosity drove me to delve deeper into this subject. After studying various academic papers and articles, I decided to implement PGO myself, benchmarking its impact to assess its true value.

My primary goal was to understand the challenges hindering PGO adoption and to identify key questions that could reveal customers' real pain points. Additionally, I aimed to explore the business value for end users. Specifically, I wanted to quantify how PGO impacts critical businesss KPIs such as conversion rates, latency, and even SLOs and SLAs.

This blog summarizes my initial findings.

**Key Takeaways:**

- This blog offers a practical guide to implementing PGO, including insights on measuring performance gains through inlining, binary size analysis, and flamegraphs.
  
- In the example code provided, my analysis revealed a notable performance gain of approximately `6.92%` in compute efficiency—an impressive result considering it’s based on a small JSON unmarshalling task. The potential savings in a production environment could be even more substantial.

- PGO can significantly boost code performance and optimize resource utilization, with the potential to reduce cloud spending by up to `14%` without any code changes.

- Many developers and SREs are missing out on potential cost savings by not leveraging PGO. You can learn from Cloudflare's [experience](https://blog.cloudflare.com/reclaiming-cpu-for-free-with-pgo) in reducing costs through PGO.

- Continuous profiling in production is essential to fully unlock the benefits of PGO.


  <img  width="1510"  alt="prom"  src="https://github.com/user-attachments/assets/cb6c174b-7d05-49ad-b464-0247e26d9c89">