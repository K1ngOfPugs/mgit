import os
import requests

# Replace 'USERNAME' with the GitHub username you want to scrape.
GITHUB_USERNAME = 'K1ngOfPugs'

# GitHub API URL for a user's repositories.
API_URL = f'https://api.github.com/users/{GITHUB_USERNAME}/repos'

def get_repos():
    """
    Fetches the list of public repositories for a given GitHub username.
    """
    try:
        response = requests.get(API_URL)
        response.raise_for_status()
        return response.json()
    except requests.RequestException as e:
        print(f"Error fetching repositories: {e}")
        return []

def download_repo_zip(repo_name, download_url, save_path):
    """
    Downloads a repository as a zip file and saves it in the specified path.
    """
    try:
        response = requests.get(download_url, stream=True)
        response.raise_for_status()
        
        with open(save_path, 'wb') as file:
            for chunk in response.iter_content(chunk_size=1024):
                if chunk:
                    file.write(chunk)
        
        print(f"Downloaded {repo_name} successfully.")
    except requests.RequestException as e:
        print(f"Error downloading {repo_name}: {e}")

def download_repos(repos):
    """
    Downloads all repositories as zip files and places them in the appropriate folder structure.
    """
    for repo in repos:
        repo_name = repo['name']
        download_url = repo['html_url'] + '/archive/refs/heads/main.zip'

        # Create folder structure: [Repo name] -> [dist]
        repo_dir = os.path.join(repo_name, 'dist')
        os.makedirs(repo_dir, exist_ok=True)
        
        # Save the zip file as [Repo name].zip inside [dist] folder
        zip_save_path = os.path.join(repo_dir, f'{repo_name}.zip')

        # Download the repo as a zip file
        download_repo_zip(repo_name, download_url, zip_save_path)

if __name__ == '__main__':
    os.Chdir(files)
    
    repos = get_repos()
    
    if repos:
        download_repos(repos)
    else:
        print("No repositories found or an error occurred.")
