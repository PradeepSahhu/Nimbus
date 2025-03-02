import axios from 'axios';
import Cookies from 'js-cookie';

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api';

export interface FileItem {
  id: string;
  name: string;
  contentType: string;
  type: string;
  size: number;
  created_at: string;
  updated_at: string;
}

class FileService {
  private getAuthHeaders() {
    const token = Cookies.get('auth_token') || Cookies.get('token');
    return token ? { 'Authorization': `Bearer ${token}` } : {};
  }

  async uploadFile(
    file: globalThis.File, 
    folderId?: string,
    onProgress?: (progressEvent: any) => void
  ): Promise<any> {
    const formData = new FormData();
    formData.append('file', file);
    
    if (folderId) {
      formData.append('folderId', folderId);
    }

    try {
      const response = await axios.post(`${API_URL}/file/upload`, formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
          ...this.getAuthHeaders()
        },
        withCredentials: true,
        onUploadProgress: onProgress,
        timeout: 60000, // 60 seconds
      });
      return response.data;
    } catch (error) {
      console.error('Upload error details:', error);
      throw error;
    }
  }

  async listFiles(): Promise<FileItem[]> {
    try {
      const response = await axios.get(`${API_URL}/file/list`, {
        headers: this.getAuthHeaders(),
        withCredentials: true,
      });
      return response.data.files || [];
    } catch (error) {
      console.error('List files error:', error);
      throw error;
    }
  }

  async deleteFile(fileId: string): Promise<void> {
    try {
      await axios.delete(`${API_URL}/file/delete/${fileId}`, {
        headers: this.getAuthHeaders(),
        withCredentials: true,
      });
    } catch (error) {
      console.error('Delete file error:', error);
      throw error;
    }
  }

  async downloadFile(fileId: string, fileName?: string): Promise<void> {
    try {
      const response = await axios({
        url: `${API_URL}/file/download/${fileId}`,
        method: 'GET',
        responseType: 'blob', 
        headers: this.getAuthHeaders(),
        withCredentials: true,
      });
  
      const url = window.URL.createObjectURL(new Blob([response.data]));
      
      let downloadFileName = fileName;
      
      if (!downloadFileName) {
        const contentDisposition = response.headers['content-disposition'];
        if (contentDisposition) {
          const filenameMatch = contentDisposition.match(/filename="(.+)"/);
          if (filenameMatch && filenameMatch.length > 1) {
            downloadFileName = filenameMatch[1];
          }
        }
      }
      
      if (!downloadFileName) {
        downloadFileName = `file-${fileId}`;
      }
      
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download', downloadFileName);
      document.body.appendChild(link);
      link.click();
      
      link.parentNode?.removeChild(link);
      window.URL.revokeObjectURL(url);
    } catch (error) {
      console.error('Download error:', error);
      throw error;
    }
  }

  getFileTypeIcon(type: string): string {
    switch (type) {
      case 'image':
        return 'image';
      case 'document':
        return 'file-text';
      case 'video':
        return 'video';
      case 'audio':
        return 'music';
      case 'archive':
        return 'archive';
      default:
        return 'file';
    }
  }

  formatFileSize(bytes: number): string {
    const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    if (bytes === 0) return '0 Byte';
    const i = Math.floor(Math.log(bytes) / Math.log(1024));
    return Math.round((bytes / Math.pow(1024, i)) * 100) / 100 + ' ' + sizes[i];
  }
}

export default new FileService();