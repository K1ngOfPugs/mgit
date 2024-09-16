import os
import json

# Directory path to scan for folders
directory_path = './files'

# Path to save the JSON output
output_json_path = './files/mgit.json'

def list_folders_in_directory(directory):
    """
    List folders in the specified directory and format the output as a list of dictionaries.
    Each dictionary contains 'name' and 'folder' keys.
    """
    folder_data = []
    
    for folder_name in os.listdir(directory):
        folder_path = os.path.join(directory, folder_name)
        if os.path.isdir(folder_path):
            folder_data.append({
                "name": folder_name,
                "folder": folder_name  # You can modify this if 'folder' should differ from 'name'
            })
    
    return folder_data

def save_to_json(data, file_path):
    """
    Save the provided data to a JSON file.
    """
    with open(file_path, 'w') as json_file:
        json.dump(data, json_file, indent=4)
        print(f"Folder data saved to {file_path}")

if __name__ == '__main__':
    # Get the folder data
    folder_data = list_folders_in_directory(directory_path)
    
    # Save the data to a JSON file
    save_to_json(folder_data, output_json_path)
