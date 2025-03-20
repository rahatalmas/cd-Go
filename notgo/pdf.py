import PyPDF2

def read_pdf(file_path):
    with open(file_path, 'rb') as file:
        reader = PyPDF2.PdfReader(file)
        text = ""
        for page in reader.pages:
            text += page.extract_text() + "\n"
    return text

if __name__ == "__main__":
    pdf_path = "example.pdf"
    extracted_text = read_pdf(pdf_path)
    print(extracted_text)


