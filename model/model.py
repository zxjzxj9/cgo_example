#! /usr/bin/env python

# see https://pytorch.org/hub/facebookresearch_pytorch-gan-zoo_dcgan/
import torch

model = torch.hub.load('facebookresearch/pytorch_GAN_zoo:hub', 'DCGAN', pretrained=True, useGPU=False)
num_images = 64
noise, _ = model.buildNoiseData(num_images)
model.netG.eval()
print(type(model.netG))
# print(noise.shape)
# with torch.no_grad():
# generated_images = model.test(noise)
# let's plot these images using torchvision and matplotlib
m: torch.jit.TracedModule = torch.jit.trace(model.netG, (noise,))
print(type(m))
torch.jit.save(m, "dcgan.pth")
# print(type(m))
# import matplotlib.pyplot as plt
# import torchvision
# plt.imshow(torchvision.utils.make_grid(generated_images).permute(1, 2, 0).cpu().numpy())