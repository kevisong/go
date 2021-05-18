# TF-IDF

## Illustration

**TF-IDF (Term Frequency-Inverse Document Frequency)**：The importance of a term is in `proportion` to `the number of appearance in the document`, and in `inverse proportion` to `the number of appearance in the corpus`.

**TF (Term Frequency)** = `# of term appeared in doc` / `total # of words in doc`

**IDF (Inverse Document Frequency)** = log( `total # of docs` / `# of docs containing the term` )

> denominator needed to be adjusted by +1 if none of the doc in corpus contains the term

**TF-IDF** = `TF * IDF`

`Cosine Similarity` = cos(( A \* B ) / ( |A| * |B| ) )
