import knorry, { type JSONData, type JSONObject, type KnorryResponseObj } from 'knorry'
import { defineKnorryOptions } from 'knorry'
import { getCredentials } from './login'
defineKnorryOptions({
    easyMode: false
})

export async function searchIcons(query: string): Promise<string[]> {
    const req = await knorry('GET', `/api/icongen/search?q=${encodeURIComponent(query)}`, null, {
        headers: {
            Authorization: `Bearer ${getCredentials()}`
        }
    }) as KnorryResponseObj

    return req.data as unknown as string[]
}

export function getAverageBackgroundColor(imgElement: HTMLImageElement): Promise<string> {
    return new Promise((resolve, reject) => {
        const canvas = document.createElement("canvas");
        const ctx = canvas.getContext("2d");

        if (!ctx) {
            reject("Could not create canvas context.");
            return;
        }

        const img = new Image();
        img.crossOrigin = "Anonymous"; // Prevent CORS issues
        img.src = imgElement.src;

        img.onload = function () {
            const width = img.width;
            const height = img.height;
            canvas.width = width;
            canvas.height = height;
            ctx.drawImage(img, 0, 0, width, height);

            const imageData = ctx.getImageData(0, 0, width, height).data;
            let r = 0, g = 0, b = 0, a = 0, count = 0;

            function processPixel(x: number, y: number) {
                const index = (y * width + x) * 4;
                const alpha = imageData[index + 3];

                if (alpha > 0) { // Ignore fully transparent pixels
                    r += imageData[index];     // Red
                    g += imageData[index + 1]; // Green
                    b += imageData[index + 2]; // Blue
                    a += alpha;                // Alpha
                    count++;
                }
            }

            // Process outermost pixels (top, bottom, left, right)
            for (let x = 0; x < width; x++) {
                processPixel(x, 0);             // Top row
                processPixel(x, height - 1);    // Bottom row
            }
            for (let y = 0; y < height; y++) {
                processPixel(0, y);             // Left column
                processPixel(width - 1, y);     // Right column
            }

            if (count === 0) {
                resolve("rgba(0, 0, 0, 0)"); // Fully transparent
            } else {
                resolve(`rgba(${Math.floor(r / count)}, ${Math.floor(g / count)}, ${Math.floor(b / count)}, ${(a / count / 255).toFixed(2)})`);
            }
        };

        img.onerror = () => reject("Error loading image.");
    });
}

// Example usage:
const imgTag = document.querySelector("img") as HTMLImageElement;